// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FilesClosure is the golang structure of table chaosplus_files_closure for DAO operations like Where/Data.
type FilesClosure struct {
	g.Meta       `orm:"table:chaosplus_files_closure, do:true"`
	AncestorId   interface{} // ancestor id
	DescendantId interface{} // descendant id
	Distance     interface{} // distance
}
