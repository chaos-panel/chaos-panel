// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileUploadingTasks is the golang structure for table file_uploading_tasks.
type FileUploadingTasks struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" ` // expired at
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // created at
	UpdatedBy uint64      `json:"updatedBy" orm:"updated_by" ` // updated by
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // updated at
}
