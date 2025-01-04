// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-17 00:12:10
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure of table logs for DAO operations like Where/Data.
type Logs struct {
	g.Meta    `orm:"table:logs, do:true"`
	Id        interface{} // ID
	Status    interface{} // 状态
	Log       interface{} // 日志
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建时间
}
