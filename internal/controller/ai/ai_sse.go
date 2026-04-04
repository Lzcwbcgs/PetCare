package ai

import (
	"encoding/json"

	"github.com/gogf/gf/v2/net/ghttp"
)

func writeSSEEvent(r *ghttp.Request, event string, payload any) {
	data, _ := json.Marshal(payload)
	r.Response.Writef("event: %s\n", event)
	r.Response.Writef("data: %s\n\n", string(data))
	r.Response.Flush()
}
