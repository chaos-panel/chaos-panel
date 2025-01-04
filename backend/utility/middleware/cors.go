package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// add cors to response
func ApiRequestDefaultCorsHandler(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
