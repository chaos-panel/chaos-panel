// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-07 00:57:54
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Settings is the golang structure of table chaosplus_settings for DAO operations like Where/Data.
type Settings struct {
	g.Meta    `orm:"table:chaosplus_settings, do:true"`
	Id        interface{} // id
	TenantId  interface{} // tenant id
	Group     interface{} // group
	Key       interface{} // key
	KeyName   interface{} // key display
	Val       interface{} // value
	ValType   interface{} // value type
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // locked at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // time updated
	DeletedBy interface{} // deleted by
	DeletedAt *gtime.Time // time deleted
}
