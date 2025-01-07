// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-07 10:37:30
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenant is the golang structure of table chaosplus_tenant for DAO operations like Where/Data.
type Tenant struct {
	g.Meta    `orm:"table:chaosplus_tenant, do:true"`
	Id        interface{} // id
	Code      interface{} // code
	Name      interface{} // name
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // locked at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // time updated
	DeletedBy interface{} // deleted by
	DeletedAt *gtime.Time // time deleted
}
