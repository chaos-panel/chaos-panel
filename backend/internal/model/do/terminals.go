// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Terminals is the golang structure of table chaosplus_terminals for DAO operations like Where/Data.
type Terminals struct {
	g.Meta    `orm:"table:chaosplus_terminals, do:true"`
	Id        interface{} // ID
	TenantId  interface{} // tenant id
	Type      interface{} // type
	Host      interface{} // host
	Port      interface{} // port
	Username  interface{} // username
	Password  interface{} // password
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // updated at
	DeletedBy interface{} // deleted by
	DeletedAt *gtime.Time // deleted at
}
