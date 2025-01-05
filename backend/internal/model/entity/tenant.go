// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure for table tenant.
type Tenant struct {
	Id          uint64      `json:"id"          orm:"id"           ` // ID
	Code        string      `json:"code"        orm:"code"         ` // 代号
	Name        string      `json:"name"        orm:"name"         ` // 名称
	TimeCreated *gtime.Time `json:"timeCreated" orm:"time_created" ` // 创建时间
	TimeUpdated *gtime.Time `json:"timeUpdated" orm:"time_updated" ` // 更新时间
}
