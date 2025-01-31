// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PluginResources is the golang structure for table plugin_resources.
type PluginResources struct {
	Id          uint64      `json:"id"          orm:"id"           ` // ID
	PluginId    uint64      `json:"pluginId"    orm:"plugin_id"    ` // plugin id
	FileType    string      `json:"fileType"    orm:"file_type"    ` // file type
	FileName    string      `json:"fileName"    orm:"file_name"    ` // file name
	FileId      uint64      `json:"fileId"      orm:"file_id"      ` // file id
	MountType   string      `json:"mountType"   orm:"mount_type"   ` // mount type
	MountPath   string      `json:"mountPath"   orm:"mount_path"   ` // mount path
	MountPolicy string      `json:"mountPolicy" orm:"mount_policy" ` // mount policy
	CreatedBy   uint64      `json:"createdBy"   orm:"created_by"   ` // created by
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // created at
	UpdatedBy   uint64      `json:"updatedBy"   orm:"updated_by"   ` // updated by
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // updated at
	DeletedBy   uint64      `json:"deletedBy"   orm:"deleted_by"   ` // deleted by
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   ` // deleted at
}
