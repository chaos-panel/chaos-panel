// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionsAdditional is the golang structure for table permissions_additional.
type PermissionsAdditional struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	RoleId    uint64      `json:"roleId"    orm:"role_id"    ` // role id
	Type      string      `json:"type"      orm:"type"       ` // type
	Value     string      `json:"value"     orm:"value"      ` // value
	Effect    string      `json:"effect"    orm:"effect"     ` // effect
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // created at
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // updated at
}
