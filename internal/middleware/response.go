package middleware

import (
	"net/http"

	"PetCare/internal/consts"
	"PetCare/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ResponseHandler wraps handler return values into the project's standard response structure.
func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 || r.Response.BytesWritten() > 0 {
		return
	}

	var (
		err     = r.GetError()
		data    = r.GetHandlerResponse()
		message = r.GetCtxVar(consts.CtxKeyResponseMessage, "success").String()
		code    = 200
	)

	if err != nil {
		message = err.Error()
		if errorCode := gerror.Code(err); errorCode != nil && errorCode != gcode.CodeNil {
			code = errorCode.Code()
		} else {
			code = http.StatusInternalServerError
		}
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		code = r.Response.Status
		message = http.StatusText(r.Response.Status)
		if message == "" {
			message = "request failed"
		}
	}

	r.Response.WriteJson(model.Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

