// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthTokens is the golang structure for table auth_tokens.
type AuthTokens struct {
	Id                    uint64      `json:"id"                    orm:"id"                       ` // ID
	TenantId              uint64      `json:"tenantId"              orm:"tenant_id"                ` // tenant id
	UserId                uint64      `json:"userId"                orm:"user_id"                  ` // user id
	ClientInfo            string      `json:"clientInfo"            orm:"client_info"              ` // client info
	AccessToken           string      `json:"accessToken"           orm:"access_token"             ` // access token
	AccessTokenJti        string      `json:"accessTokenJti"        orm:"access_token_jti"         ` // access token jti
	AccessTokenExpiredAt  *gtime.Time `json:"accessTokenExpiredAt"  orm:"access_token_expired_at"  ` // access token expires at
	RefreshToken          string      `json:"refreshToken"          orm:"refresh_token"            ` // refresh token
	RefreshTokenExpiredAt *gtime.Time `json:"refreshTokenExpiredAt" orm:"refresh_token_expired_at" ` // refresh token expires at
	CreatedBy             uint64      `json:"createdBy"             orm:"created_by"               ` // created by
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"               ` // created at
	UpdatedBy             uint64      `json:"updatedBy"             orm:"updated_by"               ` // updated by
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"               ` // updated at
	DeletedBy             uint64      `json:"deletedBy"             orm:"deleted_by"               ` // deleted by
	DeletedAt             *gtime.Time `json:"deletedAt"             orm:"deleted_at"               ` // deleted at
}
