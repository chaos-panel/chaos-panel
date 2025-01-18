// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-18 00:02:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Terminals is the golang structure for table terminals.
type Terminals struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	TenantId  uint64      `json:"tenantId"  orm:"tenant_id"  ` // tenant id
	Type      string      `json:"type"      orm:"type"       ` // type
	Host      string      `json:"host"      orm:"host"       ` // host
	Port      uint64      `json:"port"      orm:"port"       ` // port
	Username  string      `json:"username"  orm:"username"   ` // username
	Password  string      `json:"password"  orm:"password"   ` // password
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // created at
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // updated at
	DeletedBy uint64      `json:"deletedBy" orm:"deleted_by" ` // deleted by
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // deleted at
}
