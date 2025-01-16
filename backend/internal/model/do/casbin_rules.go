// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CasbinRules is the golang structure of table chaosplus_casbin_rules for DAO operations like Where/Data.
type CasbinRules struct {
	g.Meta    `orm:"table:chaosplus_casbin_rules, do:true"`
	Id        interface{} // ID
	PType     interface{} //
	V0        interface{} //
	V1        interface{} //
	V2        interface{} //
	V3        interface{} //
	V4        interface{} //
	V5        interface{} //
	CreatedAt *gtime.Time //
}
