// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-06 00:06:52
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerId is the golang structure for table worker_id.
type WorkerId struct {
	Id          uint        `json:"id"          orm:"id"           ` // id
	HostInfo    string      `json:"hostInfo"    orm:"host_info"    ` // host info
	HostTag     string      `json:"hostTag"     orm:"host_tag"     ` // host tag
	TimeCreated *gtime.Time `json:"timeCreated" orm:"time_created" ` // time created
	TimeUpdated *gtime.Time `json:"timeUpdated" orm:"time_updated" ` // time updated
}
