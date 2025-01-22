// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthClients is the golang structure for table auth_clients.
type AuthClients struct {
	Id           uint64      `json:"id"           orm:"id"            ` // ID
	TenantId     uint64      `json:"tenantId"     orm:"tenant_id"     ` // tenant id
	ClientId     string      `json:"clientId"     orm:"client_id"     ` // client id
	ClientSecret string      `json:"clientSecret" orm:"client_secret" ` // client secret
	GrantTypes   string      `json:"grantTypes"   orm:"grant_types"   ` // grant types
	Scopes       string      `json:"scopes"       orm:"scopes"        ` // scopes
	RedirectUri  string      `json:"redirectUri"  orm:"redirect_uri"  ` // redirect uri
	CreatedBy    uint64      `json:"createdBy"    orm:"created_by"    ` // created by
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` // created at
	UpdatedBy    uint64      `json:"updatedBy"    orm:"updated_by"    ` // updated by
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` // updated at
	DeletedBy    uint64      `json:"deletedBy"    orm:"deleted_by"    ` // deleted by
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    ` // deleted at
}
