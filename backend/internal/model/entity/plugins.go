// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Plugins is the golang structure for table plugins.
type Plugins struct {
	Id          uint64      `json:"id"          orm:"id"           ` // ID
	TenantId    uint64      `json:"tenantId"    orm:"tenant_id"    ` // tenant id
	RepoId      uint64      `json:"repoId"      orm:"repo_id"      ` // repo id
	Kind        string      `json:"kind"        orm:"kind"         ` // kind
	Type        string      `json:"type"        orm:"type"         ` // type
	Name        string      `json:"name"        orm:"name"         ` // name
	Author      string      `json:"author"      orm:"author"       ` // author
	VersionCode int64       `json:"versionCode" orm:"version_code" ` // version code
	VersionName string      `json:"versionName" orm:"version_name" ` // version name
	Desc        string      `json:"desc"        orm:"desc"         ` // desc
	CreatedBy   uint64      `json:"createdBy"   orm:"created_by"   ` // created by
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // created at
	UpdatedBy   uint64      `json:"updatedBy"   orm:"updated_by"   ` // updated by
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // updated at
	DeletedBy   uint64      `json:"deletedBy"   orm:"deleted_by"   ` // deleted by
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   ` // deleted at
}
