package middleware

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chaos-plus/chaos-plus/internal/consts"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/shortid"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func ApiRequestAuthHandler(r *ghttp.Request) {
	if !IsWhiteList(r) && VerifyToken(r) == nil {
		r.Response.WriteStatus(http.StatusUnauthorized, &ghttp.DefaultHandlerResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    nil,
		})
	} else {
		r.Middleware.Next()
	}
}

func GetUserId(r *ghttp.Request) int64 {
	tok := VerifyToken(r)
	if tok == nil {
		return 0
	} else {
		return gconv.Int64(tok.Subject())
	}
}

func getToken(r *ghttp.Request) string {
	token1 := r.GetHeader("Authorization")
	if len(token1) > 0 {
		return token1
	}
	token2 := r.GetHeader("Token")
	if len(token2) > 0 {
		return token2
	}
	token3 := r.GetParam("Token")
	if token3 == nil {
		return ""
	} else {
		return token3.String()
	}
}

func IsWhiteList(r *ghttp.Request) bool {
	for i := range consts.AUTH_WHITE_LIST {
		white := consts.AUTH_WHITE_LIST[i]
		if !strings.Contains(white, ":") {
			continue
		}
		split := strings.Split(white, ":")
		method := split[0]
		uri := split[1]
		if strings.EqualFold(r.Method, method) && strings.Contains(r.RequestURI, uri) {
			return true
		}
	}
	return false
}

func VerifyToken(r *ghttp.Request) jwt.Token {
	token := ParseToken(r)
	if token == nil {
		return nil
	}
	if r.Host != token.Issuer() {
		fmt.Printf("token issuer not match %s != %s\n", r.Host, token.Issuer())
		return nil
	}
	if gtime.Now().Time.Before(token.NotBefore()) {
		fmt.Printf("token can use not before %v<%v\n", gtime.Now().Time, token.NotBefore())
		return nil
	}
	return token
}

func ParseToken(r *ghttp.Request) jwt.Token {
	typ := "Bearer"
	token := getToken(r)
	token = strings.Trim(token, " ")
	if token == "" {
		return nil
	}
	if strings.Contains(token, " ") {
		split := strings.Split(token, " ")
		typ = split[0]
		token = split[1]
	}
	if typ != "Bearer" {
		return nil
	}
	parse, err := JwtParse([]byte(token))
	if err != nil {
		return nil
	}
	return parse
}

func JwtParse(token []byte) (jwt.Token, error) {
	block, _ := pem.Decode(([]byte)(consts.JWT_RSA_PUB_KEY))
	pubkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("failed to get public key: %s\n", err)
		return nil, err
	}
	tok, err := jwt.Parse(token, jwt.WithKey(jwa.RS256, pubkey))
	if err != nil {
		fmt.Printf("failed to verify JWT: %s\n", err)
		return nil, err
	}
	return tok, nil
}

func JwtSign(r *ghttp.Request, userId string, clientId string, validity time.Duration, claims g.Map) (jwt.Token, []byte, error) {
	id := service.Guid().NextId()
	sid := shortid.NewShortId(id).Short36
	now := gtime.Now().Time
	builder := jwt.NewBuilder().
		Issuer(r.Host).                // 必须, 发行机构Issuer, 大小写敏感的URL, 不能包含query参数
		Subject(userId).               // 必须, 用户身份Subject, Issuer为End-User分配的唯一标识符, 大小写敏感不超过255 ASCII自符
		Audience([]string{clientId}).  // 必须, 特别的身份Audience, 必须包含OAuth2的client_id, 大小写敏感的字符串/数组
		Expiration(now.Add(validity)). // 必须, iat到期时间Expire, 参数要求当前时间在该时间之前, 通常可以时钟偏差几分钟, unix时间戳
		IssuedAt(now).                 // 必须, JWT颁发时间Issuer at time, unix时间戳
		NotBefore(now).                // 必须, JWT生效时间 Not Before, unix时间戳
		JwtID(sid)                     // 防重放攻击
	for k, v := range claims {
		builder.Claim(k, v)
	}
	tok, err := builder.Build()
	if err != nil {
		fmt.Printf("failed to build token: %s\n", err)
		return nil, nil, err
	}
	if err := json.NewEncoder(os.Stdout).Encode(tok); err != nil {
		fmt.Printf("failed to encode to JSON: %s\n", err)
		return nil, nil, err
	}

	// Parse, serialize, slice and dice JWKs!
	block, _ := pem.Decode(([]byte)(consts.JWT_RSA_PRI_KEY))
	privkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("failed to parse JWK: %s\n", err)
		return nil, nil, err
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privkey))
	if err != nil {
		fmt.Printf("failed to sign token: %s\n", err)
		return nil, nil, err
	}
	return tok, signed, nil
}
