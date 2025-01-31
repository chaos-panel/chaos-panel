// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Organizations is the golang structure for table organizations.
type Organizations struct {
	Id          uint64      `json:"id"          orm:"id"          ` // ID
	TenantId    uint64      `json:"tenantId"    orm:"tenant_id"   ` // tenant id
	Name        string      `json:"name"        orm:"name"        ` // name
	Code        string      `json:"code"        orm:"code"        ` // code
	Logo        string      `json:"logo"        orm:"logo"        ` // logo
	Email       string      `json:"email"       orm:"email"       ` // email
	Location    string      `json:"location"    orm:"location"    ` // location
	Website     string      `json:"website"     orm:"website"     ` // website
	Description string      `json:"description" orm:"description" ` // description
	Type        string      `json:"type"        orm:"type"        ` // type
	CreatedBy   uint64      `json:"createdBy"   orm:"created_by"  ` // created by
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  ` // created at
	UpdatedBy   uint64      `json:"updatedBy"   orm:"updated_by"  ` // updated by
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  ` // updated at
	DeletedBy   uint64      `json:"deletedBy"   orm:"deleted_by"  ` // deleted by
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  ` // deleted at
}
