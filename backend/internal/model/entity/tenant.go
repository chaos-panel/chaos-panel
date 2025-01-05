// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-06 00:06:52
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure for table tenant.
type Tenant struct {
	Id          uint64      `json:"id"          orm:"id"           ` // id
	Code        string      `json:"code"        orm:"code"         ` // code
	Name        string      `json:"name"        orm:"name"         ` // name
	TimeCreated *gtime.Time `json:"timeCreated" orm:"time_created" ` // time created
	TimeUpdated *gtime.Time `json:"timeUpdated" orm:"time_updated" ` // time updated
}
