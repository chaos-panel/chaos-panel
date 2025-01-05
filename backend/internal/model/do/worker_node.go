// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerNode is the golang structure of table chaosplus_worker_node for DAO operations like Where/Data.
type WorkerNode struct {
	g.Meta      `orm:"table:chaosplus_worker_node, do:true"`
	Id          interface{} // 节点id
	Port        interface{} // 端口
	OsName      interface{} // 系统名称
	OsArch      interface{} // 系统平台
	OsVersion   interface{} // 系统版本
	HostName    interface{} // 主机名称
	Docker      interface{} // IP地址
	Ip          interface{} // IP地址
	Mac         interface{} // MAC地址
	Tag         interface{} // 节点标签
	Desc        interface{} // 节点描述
	TimeCreated *gtime.Time // 创建时间
	TimeUpdated *gtime.Time // 更新时间
}
