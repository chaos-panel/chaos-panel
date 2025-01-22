// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	Id            uint64      `json:"id"            orm:"id"             ` // ID
	TenantId      uint64      `json:"tenantId"      orm:"tenant_id"      ` // tenant id
	Name          string      `json:"name"          orm:"name"           ` // name
	Desc          string      `json:"desc"          orm:"desc"           ` // description
	Path          string      `json:"path"          orm:"path"           ` // path
	PathDepth     uint64      `json:"pathDepth"     orm:"path_depth"     ` // path depth
	PathMd5       string      `json:"pathMd5"       orm:"path_md5"       ` // path md5
	PathSha1      string      `json:"pathSha1"      orm:"path_sha1"      ` // path sha1
	PathSha256    string      `json:"pathSha256"    orm:"path_sha256"    ` // path sha256
	PathSha512    string      `json:"pathSha512"    orm:"path_sha512"    ` // path sha512
	ChildrenCount uint64      `json:"childrenCount" orm:"children_count" ` // children node count
	CreatedBy     uint64      `json:"createdBy"     orm:"created_by"     ` // created by
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     ` // created at
	UpdatedBy     uint64      `json:"updatedBy"     orm:"updated_by"     ` // updated by
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     ` // updated at
	DeletedBy     uint64      `json:"deletedBy"     orm:"deleted_by"     ` // deleted by
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     ` // deleted at
}
