// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id           uint64      `json:"id"           orm:"id"             ` // ID
	TenantId     uint64      `json:"tenantId"     orm:"tenant_id"      ` // tenant id
	Username     string      `json:"username"     orm:"username"       ` // username
	Firstname    string      `json:"firstname"    orm:"firstname"      ` // firstname
	Lastname     string      `json:"lastname"     orm:"lastname"       ` // lastname
	Nickname     string      `json:"nickname"     orm:"nickname"       ` // nickname
	Password     string      `json:"password"     orm:"password"       ` // password
	Phone        string      `json:"phone"        orm:"phone"          ` // phone
	Email        string      `json:"email"        orm:"email"          ` // email
	AvatarFileId uint64      `json:"avatarFileId" orm:"avatar_file_id" ` // avatar file id
	WxUnionId    string      `json:"wxUnionId"    orm:"wx_union_id"    ` // wx union id
	WxmaOpenId   string      `json:"wxmaOpenId"   orm:"wxma_open_id"   ` // wxma open id
	WxmpOpenId   string      `json:"wxmpOpenId"   orm:"wxmp_open_id"   ` // wxmp open id
	GithubUserId string      `json:"githubUserId" orm:"github_user_id" ` // github user id
	LockedBy     uint64      `json:"lockedBy"     orm:"locked_by"      ` // locked by
	LockedAt     *gtime.Time `json:"lockedAt"     orm:"locked_at"      ` // locked at
	CreatedBy    uint64      `json:"createdBy"    orm:"created_by"     ` // created by
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     ` // created at
	UpdatedBy    uint64      `json:"updatedBy"    orm:"updated_by"     ` // updated by
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     ` // updated at
	DeletedBy    uint64      `json:"deletedBy"    orm:"deleted_by"     ` // deleted by
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"     ` // deleted at
}
