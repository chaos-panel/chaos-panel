// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RolesClosure is the golang structure of table chaosplus_roles_closure for DAO operations like Where/Data.
type RolesClosure struct {
	g.Meta       `orm:"table:chaosplus_roles_closure, do:true"`
	AncestorId   interface{} // ancestor id
	DescendantId interface{} // descendant id
	Distance     interface{} // distance
}
