package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func generateToken(ctx context.Context, claims AuthClaims) (string, error) {
	headerBytes, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		return "", gerror.Wrap(err, "生成 token 头失败")
	}

	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", gerror.Wrap(err, "生成 token 载荷失败")
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
		return nil, gerror.NewCode(gcode.New(401, "", nil), "token 格式不正确")
	}

	signingContent := parts[0] + "." + parts[1]
	if signToken(signingContent, jwtSecret(ctx)) != parts[2] {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "token 签名无效")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, gerror.Wrap(err, "解析 token 载荷失败")
	}

	var claims AuthClaims
	if err = json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, gerror.Wrap(err, "解析 token 失败")
	}
	if claims.Issuer != "" && claims.Issuer != g.Cfg().MustGet(ctx, "auth.jwtIssuer", "petcare").String() {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "token 签发方无效")
	}
	if claims.Exp <= time.Now().Unix() {
		return nil, gerror.NewCode(gcode.New(401, "", nil), "token 已过期")
	}
	return &claims, nil
}

func signToken(content string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(content))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func jwtSecret(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "auth.jwtSecret", "petcare-dev-secret").String()
}
