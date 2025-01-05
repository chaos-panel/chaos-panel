// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-06 00:06:52
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure for table logs.
type Logs struct {
	Id        uint64      `json:"id"        orm:"id"         ` // id
	Status    string      `json:"status"    orm:"status"     ` // status
	Log       string      `json:"log"       orm:"log"        ` // log
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // created by
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // created at
}
