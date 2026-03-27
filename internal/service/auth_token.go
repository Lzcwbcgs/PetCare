package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"PetCare/internal/consts"
)

var revokedTokenStore sync.Map

func generateToken(ctx context.Context, claims AuthClaims) (string, error) {
	headerBytes, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		return "", consts.WrapInternalError(err, "生成 token 头失败")
	}

	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", consts.WrapInternalError(err, "生成 token 载荷失败")
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(headerBytes)
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	signingContent := encodedHeader + "." + encodedPayload
	signature := signToken(signingContent, jwtSecret(ctx))
	return signingContent + "." + signature, nil
}

func parseToken(ctx context.Context, token string) (*AuthClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, consts.NewUnauthorizedError("token 格式不正确")
	}

	signingContent := parts[0] + "." + parts[1]
	if signToken(signingContent, jwtSecret(ctx)) != parts[2] {
		return nil, consts.NewUnauthorizedError("token 签名无效")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, consts.WrapInternalError(err, "解析 token 载荷失败")
	}

	var claims AuthClaims
	if err = json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, consts.WrapInternalError(err, "解析 token 失败")
	}
	if claims.UserID <= 0 || claims.Username == "" {
		return nil, consts.NewUnauthorizedError("token 载荷无效")
	}
	if !IsSupportedRole(claims.Role) {
		return nil, consts.NewUnauthorizedError("token 角色无效")
	}
	if claims.Issuer != "" && claims.Issuer != authConfigString(ctx, "auth.jwtIssuer", "petcare") {
		return nil, consts.NewUnauthorizedError("token 签发方无效")
	}
	if claims.Exp <= time.Now().Unix() {
		return nil, consts.NewUnauthorizedError("token 已过期")
	}
	return &claims, nil
}

func signToken(content string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(content))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func jwtSecret(ctx context.Context) string {
	return authConfigString(ctx, "auth.jwtSecret", "petcare-dev-secret")
}

func revokeToken(token string, ttl time.Duration) {
	if token == "" {
		return
	}
	if ttl <= 0 {
		ttl = time.Second
	}
	expireAt := time.Now().Add(ttl)
	revokedTokenStore.Store(token, expireAt)
}

func isTokenRevoked(token string) bool {
	if token == "" {
		return false
	}
	value, ok := revokedTokenStore.Load(token)
	if !ok {
		return false
	}
	expireAt, ok := value.(time.Time)
	if !ok {
		revokedTokenStore.Delete(token)
		return false
	}
	if time.Now().After(expireAt) {
		revokedTokenStore.Delete(token)
		return false
	}
	return true
}
