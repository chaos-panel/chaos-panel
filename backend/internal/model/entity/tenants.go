// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-11 20:17:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenants is the golang structure for table tenants.
type Tenants struct {
	Id        uint64      `json:"id"        orm:"id"         ` // id
	Code      string      `json:"code"      orm:"code"       ` // code
	Name      string      `json:"name"      orm:"name"       ` // name
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // locked at
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // time updated
	DeletedBy uint64      `json:"deletedBy" orm:"deleted_by" ` // deleted by
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // time deleted
}
