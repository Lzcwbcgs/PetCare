package consts

import (
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ErrorCodeBadRequest   = gcode.New(400, "Bad Request", nil)
	ErrorCodeUnauthorized = gcode.New(401, "Unauthorized", nil)
	ErrorCodeForbidden    = gcode.New(403, "Forbidden", nil)
	ErrorCodeNotFound     = gcode.New(404, "Not Found", nil)
	ErrorCodeConflict     = gcode.New(409, "Conflict", nil)
	ErrorCodeInternal     = gcode.New(500, "Internal Server Error", nil)
)

const (
	ErrorMessageBadRequest   = "请求参数错误"
	ErrorMessageUnauthorized = "未登录或 token 无效"
	ErrorMessageForbidden    = "无权限访问"
	ErrorMessageNotFound     = "资源不存在"
	ErrorMessageConflict     = "业务冲突"
	ErrorMessageInternal     = "服务器内部错误"
)

func NewBadRequestError(message string) error {
	return gerror.NewCode(ErrorCodeBadRequest, fallbackErrorMessage(message, ErrorMessageBadRequest))
}

func NewUnauthorizedError(message string) error {
	return gerror.NewCode(ErrorCodeUnauthorized, fallbackErrorMessage(message, ErrorMessageUnauthorized))
}

func NewForbiddenError(message string) error {
	return gerror.NewCode(ErrorCodeForbidden, fallbackErrorMessage(message, ErrorMessageForbidden))
}

func NewNotFoundError(message string) error {
	return gerror.NewCode(ErrorCodeNotFound, fallbackErrorMessage(message, ErrorMessageNotFound))
}

func NewConflictError(message string) error {
	return gerror.NewCode(ErrorCodeConflict, fallbackErrorMessage(message, ErrorMessageConflict))
}

func WrapInternalError(err error, message string) error {
	return gerror.WrapCode(ErrorCodeInternal, err, fallbackErrorMessage(message, ErrorMessageInternal))
}

func fallbackErrorMessage(message string, fallback string) string {
	if strings.TrimSpace(message) == "" {
		return fallback
	}
	return message
}
