// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-11 20:17:21
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerIds is the golang structure for table worker_ids.
type WorkerIds struct {
	Id        uint        `json:"id"        orm:"id"         ` // id
	HostInfo  string      `json:"hostInfo"  orm:"host_info"  ` // host info
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" ` // expired at
	CreatedBy string      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // locked at
	UpdatedBy string      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // time updated
}
