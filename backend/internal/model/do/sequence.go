// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Sequence is the golang structure of table chaosplus_sequence for DAO operations like Where/Data.
type Sequence struct {
	g.Meta   `orm:"table:chaosplus_sequence, do:true"`
	Id       interface{} // ID
	Type     interface{} // type
	Sequence interface{} // sequence
}
