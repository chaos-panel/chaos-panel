// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthClients is the golang structure of table chaosplus_auth_clients for DAO operations like Where/Data.
type AuthClients struct {
	g.Meta       `orm:"table:chaosplus_auth_clients, do:true"`
	Id           interface{} // ID
	TenantId     interface{} // tenant id
	ClientId     interface{} // client id
	ClientSecret interface{} // client secret
	GrantTypes   interface{} // grant types
	Scopes       interface{} // scopes
	RedirectUri  interface{} // redirect uri
	CreatedBy    interface{} // created by
	CreatedAt    *gtime.Time // created at
	UpdatedBy    interface{} // updated by
	UpdatedAt    *gtime.Time // updated at
	DeletedBy    interface{} // deleted by
	DeletedAt    *gtime.Time // deleted at
}
