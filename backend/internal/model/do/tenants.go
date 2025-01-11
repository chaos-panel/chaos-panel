// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-11 20:17:21
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tenants is the golang structure of table chaosplus_tenants for DAO operations like Where/Data.
type Tenants struct {
	g.Meta    `orm:"table:chaosplus_tenants, do:true"`
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
