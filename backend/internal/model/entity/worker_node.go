// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerNode is the golang structure for table worker_node.
type WorkerNode struct {
	Id          uint        `json:"id"          orm:"id"           ` // 节点id
	Port        uint        `json:"port"        orm:"port"         ` // 端口
	OsName      string      `json:"osName"      orm:"os_name"      ` // 系统名称
	OsArch      string      `json:"osArch"      orm:"os_arch"      ` // 系统平台
	OsVersion   string      `json:"osVersion"   orm:"os_version"   ` // 系统版本
	HostName    string      `json:"hostName"    orm:"host_name"    ` // 主机名称
	Docker      int         `json:"docker"      orm:"docker"       ` // IP地址
	Ip          string      `json:"ip"          orm:"ip"           ` // IP地址
	Mac         string      `json:"mac"         orm:"mac"          ` // MAC地址
	Tag         string      `json:"tag"         orm:"tag"          ` // 节点标签
	Desc        string      `json:"desc"        orm:"desc"         ` // 节点描述
	TimeCreated *gtime.Time `json:"timeCreated" orm:"time_created" ` // 创建时间
	TimeUpdated *gtime.Time `json:"timeUpdated" orm:"time_updated" ` // 更新时间
}
