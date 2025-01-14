// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-14 17:54:17
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChangeLogs is the golang structure for table change_logs.
type ChangeLogs struct {
	Id        uint64      `json:"id"        orm:"id"         ` // id
	UnionId   uint64      `json:"unionId"   orm:"union_id"   ` // union id
	Snapshot  string      `json:"snapshot"  orm:"snapshot"   ` // snapshot
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // locked at
	DeletedBy uint64      `json:"deletedBy" orm:"deleted_by" ` // deleted by
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // time deleted
}
