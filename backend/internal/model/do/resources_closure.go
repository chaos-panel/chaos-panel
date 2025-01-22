// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ResourcesClosure is the golang structure of table chaosplus_resources_closure for DAO operations like Where/Data.
type ResourcesClosure struct {
	g.Meta       `orm:"table:chaosplus_resources_closure, do:true"`
	AncestorId   interface{} // ancestor id
	DescendantId interface{} // descendant id
	Distance     interface{} // distance
}
