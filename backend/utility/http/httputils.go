package httputils

import (
	"strings"

	"github.com/chaos-plus/chaos-plus/utility/geoip"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func GetToken(r *ghttp.Request) (string, string) {
	token := r.Header.Get("Authorization")
	if token == "" {
		token = r.Header.Get("X-Token")
	}
	if token == "" {
		token = r.GetQuery("token").String()
	}
	idx := strings.Index(token, " ")
	if idx > 0 && idx+1 < len(token) {
		return token[0:idx], token[idx+1:]
	}
	return "", token

}

func GetTenantId(r *ghttp.Request) string {
	return r.Header.Get("X-Tenant-Id")
}

func GetUserAgent(r *ghttp.Request) string {
	return r.Header.Get("User-Agent")
}

func GetClientIP(r *ghttp.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
		if ip != "" {
			ips := strings.Split(ip, ",")
			ip = strings.TrimSpace(ips[0])
		}
	}
	if ip == "" {
		ip = r.Header.Get("Proxy-Client-IP")
	}
	if ip == "" {
		ip = r.Header.Get("WL-Proxy-Client-IP")
	}
	if ip == "" {
		ip = r.RemoteAddr
		if idx := strings.Index(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
	}
	if ip == "" {
		return "unknown"
	}
	return ip
}

func GetClientIpLocation(r *ghttp.Request) string {
	ip := GetClientIP(r)
	return geoip.GetLocation(ip)
}

func GetClientInfo(r *ghttp.Request) *gjson.Json {
	ua := GetUserAgent(r)
	ip := GetClientIP(r)
	location := geoip.GetLocation(ip)

	info := gjson.New(make(map[string]interface{}))
	info.Set("ua", ua)
	info.Set("ip", ip)
	info.Set("location", location)
	return info
}
