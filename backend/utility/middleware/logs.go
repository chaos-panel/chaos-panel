package middleware

import (
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/service"
	httputils "github.com/chaos-plus/chaos-plus/utility/http"
)

type LogContent struct {
	Type     string      `json:"type"      dc:"log type"`
	Request  interface{} `json:"request"   dc:"operation command"`
	Response interface{} `json:"response"  dc:"operation result"`
	Error    interface{} `json:"error"  dc:"operation result"`
}

func ApiRequestLogsHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// skip GET
	if strings.ToUpper(r.Method) == "GET" {
		return
	}
	// skip OPTIONS
	if strings.ToUpper(r.Method) == "OPTIONS" {
		return
	}
	requestMap := g.Map{
		"method":  r.Request.Method,
		"uri":     r.Request.RequestURI,
		"headers": r.Request.Header,
		"body":    r.GetBodyString(),
	}

	responseMap := g.Map{
		"body": r.GetHandlerResponse(),
	}

	var errorStatus string
	var errorMessage string
	if r.GetError() != nil {
		errorStatus = "ERROR"
		errorMessage = r.GetError().Error()
	} else {
		errorStatus = ""
		errorMessage = ""
	}

	content := &LogContent{
		Type:     "Http",
		Request:  requestMap,
		Response: responseMap,
		Error:    errorMessage,
	}
	json, err := gjson.EncodeString(content)
	if err != nil {
		g.Log().Error(r.GetCtx(), err)
		return
	}
	service.Logs().Create(r.GetCtx(), &do.Logs{
		Id:         service.Guid().NextId(),
		TenantId:   httputils.GetTenantId(r),
		ClientType: "http api middleware",
		ClientInfo: httputils.GetClientInfo(r),
		Remark:     errorStatus,
		Log:        json,
		CreatedBy:  GetUserId(r), // TODO
		CreatedAt:  gtime.Now(),
	})
}
