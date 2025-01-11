// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-11 20:17:21
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TSequence is the golang structure of table t_sequence for DAO operations like Where/Data.
type TSequence struct {
	g.Meta   `orm:"table:t_sequence, do:true"`
	Id       interface{} // ID
	Type     interface{} // type
	Sequence interface{} // sequence
}
