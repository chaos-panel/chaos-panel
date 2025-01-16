// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Labels is the golang structure of table chaosplus_labels for DAO operations like Where/Data.
type Labels struct {
	g.Meta     `orm:"table:chaosplus_labels, do:true"`
	LabelId    interface{} // label id
	UnionId    interface{} // union id
	LabelIndex interface{} // label index
	LabelName  interface{} // label name
	LabelColor interface{} // label color
	CreatedBy  interface{} // created by
	CreatedAt  *gtime.Time // locked at
}
