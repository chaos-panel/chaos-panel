// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Settings is the golang structure for table settings.
type Settings struct {
	Id        uint64      `json:"id"        orm:"id"         ` // id
	TenantId  uint64      `json:"tenantId"  orm:"tenant_id"  ` // tenant id
	Group     string      `json:"group"     orm:"group"      ` // group
	Key       string      `json:"key"       orm:"key"        ` // key
	KeyName   string      `json:"keyName"   orm:"key_name"   ` // key display
	Val       string      `json:"val"       orm:"val"        ` // value
	ValType   string      `json:"valType"   orm:"val_type"   ` // value type
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // locked at
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // time updated
	DeletedBy uint64      `json:"deletedBy" orm:"deleted_by" ` // deleted by
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // time deleted
}
