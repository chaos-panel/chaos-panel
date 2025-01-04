// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-17 00:12:10
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure for table logs.
type Logs struct {
	Id        uint64      `json:"id"        orm:"id"         ` // ID
	Status    string      `json:"status"    orm:"status"     ` // 状态
	Log       string      `json:"log"       orm:"log"        ` // 日志
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // 创建者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
}
