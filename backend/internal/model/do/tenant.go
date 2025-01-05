// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure of table chaosplus_tenant for DAO operations like Where/Data.
type Tenant struct {
	g.Meta      `orm:"table:chaosplus_tenant, do:true"`
	Id          interface{} // ID
	Code        interface{} // 代号
	Name        interface{} // 名称
	TimeCreated *gtime.Time // 创建时间
	TimeUpdated *gtime.Time // 更新时间
}
