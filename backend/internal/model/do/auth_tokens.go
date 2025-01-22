// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthTokens is the golang structure of table chaosplus_auth_tokens for DAO operations like Where/Data.
type AuthTokens struct {
	g.Meta                `orm:"table:chaosplus_auth_tokens, do:true"`
	Id                    interface{} // ID
	TenantId              interface{} // tenant id
	UserId                interface{} // user id
	ClientInfo            interface{} // client info
	AccessToken           interface{} // access token
	AccessTokenJti        interface{} // access token jti
	AccessTokenExpiredAt  *gtime.Time // access token expires at
	RefreshToken          interface{} // refresh token
	RefreshTokenExpiredAt *gtime.Time // refresh token expires at
	CreatedBy             interface{} // created by
	CreatedAt             *gtime.Time // created at
	UpdatedBy             interface{} // updated by
	UpdatedAt             *gtime.Time // updated at
	DeletedBy             interface{} // deleted by
	DeletedAt             *gtime.Time // deleted at
}
