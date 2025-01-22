// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Organizations is the golang structure of table chaosplus_organizations for DAO operations like Where/Data.
type Organizations struct {
	g.Meta      `orm:"table:chaosplus_organizations, do:true"`
	Id          interface{} // ID
	TenantId    interface{} // tenant id
	Name        interface{} // name
	Code        interface{} // code
	Logo        interface{} // logo
	Email       interface{} // email
	Location    interface{} // location
	Website     interface{} // website
	Description interface{} // description
	Type        interface{} // type
	CreatedBy   interface{} // created by
	CreatedAt   *gtime.Time // created at
	UpdatedBy   interface{} // updated by
	UpdatedAt   *gtime.Time // updated at
	DeletedBy   interface{} // deleted by
	DeletedAt   *gtime.Time // deleted at
}
