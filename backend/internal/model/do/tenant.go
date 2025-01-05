// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-06 00:06:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure of table chaosplus_tenant for DAO operations like Where/Data.
type Tenant struct {
	g.Meta      `orm:"table:chaosplus_tenant, do:true"`
	Id          interface{} // id
	Code        interface{} // code
	Name        interface{} // name
	TimeCreated *gtime.Time // time created
	TimeUpdated *gtime.Time // time updated
}
