// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersRoles is the golang structure for table users_roles.
type UsersRoles struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	UserId    uint64      `json:"userId"    orm:"user_id"    ` // user id
	RoleId    uint64      `json:"roleId"    orm:"role_id"    ` // role id
	ValidedAt *gtime.Time `json:"validedAt" orm:"valided_at" ` // valided at
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" ` // expired at
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // created at
}
