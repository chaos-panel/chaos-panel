// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure for table permissions.
type Permissions struct {
	Id         uint64      `json:"id"         orm:"id"          ` // ID
	RoleId     uint64      `json:"roleId"     orm:"role_id"     ` // role id
	ResourceId uint64      `json:"resourceId" orm:"resource_id" ` // resource id
	Action     string      `json:"action"     orm:"action"      ` // action
	Effect     string      `json:"effect"     orm:"effect"      ` // effect
	CreatedBy  uint64      `json:"createdBy"  orm:"created_by"  ` // created by
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // created at
	UpdatedBy  uint64      `json:"updatedBy"  orm:"updated_by"  ` // updated by
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` // updated at
}
