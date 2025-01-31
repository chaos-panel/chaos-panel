// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure for table files.
type Files struct {
	Id                 uint64      `json:"id"                 orm:"id"                    ` // ID
	BucketId           uint64      `json:"bucketId"           orm:"bucket_id"             ` // bucket id
	IsDir              uint        `json:"isDir"              orm:"is_dir"                ` // is dir
	IsSystem           uint        `json:"isSystem"           orm:"is_system"             ` // is system file
	IsHidden           uint        `json:"isHidden"           orm:"is_hidden"             ` // is hidden file
	Size               uint64      `json:"size"               orm:"size"                  ` // size
	Name               string      `json:"name"               orm:"name"                  ` // name
	Path               string      `json:"path"               orm:"path"                  ` // path
	PathDepth          uint64      `json:"pathDepth"          orm:"path_depth"            ` // path depth
	PathMd5            string      `json:"pathMd5"            orm:"path_md5"              ` // path md5
	PathSha1           string      `json:"pathSha1"           orm:"path_sha1"             ` // path sha1
	PathSha256         string      `json:"pathSha256"         orm:"path_sha256"           ` // path sha256
	PathSha512         string      `json:"pathSha512"         orm:"path_sha512"           ` // path sha512
	Ext                string      `json:"ext"                orm:"ext"                   ` // extension name
	Type               string      `json:"type"               orm:"type"                  ` // type
	Subtype            string      `json:"subtype"            orm:"subtype"               ` // subtype
	FileId             uint64      `json:"fileId"             orm:"file_id"               ` // physical file id
	Desc               string      `json:"desc"               orm:"desc"                  ` // description
	SubFileCount       uint64      `json:"subFileCount"       orm:"sub_file_count"        ` // sub file node count
	SubFileSystemCount uint64      `json:"subFileSystemCount" orm:"sub_file_system_count" ` // sub file system node count
	SubFileHiddenCount uint64      `json:"subFileHiddenCount" orm:"sub_file_hidden_count" ` // sub file hidden node count
	SubDirCount        uint64      `json:"subDirCount"        orm:"sub_dir_count"         ` // sub dir node count
	SubDirSystemCount  uint64      `json:"subDirSystemCount"  orm:"sub_dir_system_count"  ` // sub dir system node count
	SubDirHiddenCount  uint64      `json:"subDirHiddenCount"  orm:"sub_dir_hidden_count"  ` // sub dir hidden node count
	CreatedBy          uint64      `json:"createdBy"          orm:"created_by"            ` // created by
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"            ` // created at
	UpdatedBy          uint64      `json:"updatedBy"          orm:"updated_by"            ` // updated by
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"            ` // updated at
	DeletedBy          uint64      `json:"deletedBy"          orm:"deleted_by"            ` // deleted by
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"            ` // deleted at
}
