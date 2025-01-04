package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// disable cache
func ApiRequestNoCacheHandler(r *ghttp.Request) {
	r.Response.Header().Add("Cache-control", "no-store")
	r.Middleware.Next()
}
