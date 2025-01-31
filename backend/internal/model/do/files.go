// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure of table chaosplus_files for DAO operations like Where/Data.
type Files struct {
	g.Meta             `orm:"table:chaosplus_files, do:true"`
	Id                 interface{} // ID
	BucketId           interface{} // bucket id
	IsDir              interface{} // is dir
	IsSystem           interface{} // is system file
	IsHidden           interface{} // is hidden file
	Size               interface{} // size
	Name               interface{} // name
	Path               interface{} // path
	PathDepth          interface{} // path depth
	PathMd5            interface{} // path md5
	PathSha1           interface{} // path sha1
	PathSha256         interface{} // path sha256
	PathSha512         interface{} // path sha512
	Ext                interface{} // extension name
	Type               interface{} // type
	Subtype            interface{} // subtype
	FileId             interface{} // physical file id
	Desc               interface{} // description
	SubFileCount       interface{} // sub file node count
	SubFileSystemCount interface{} // sub file system node count
	SubFileHiddenCount interface{} // sub file hidden node count
	SubDirCount        interface{} // sub dir node count
	SubDirSystemCount  interface{} // sub dir system node count
	SubDirHiddenCount  interface{} // sub dir hidden node count
	CreatedBy          interface{} // created by
	CreatedAt          *gtime.Time // created at
	UpdatedBy          interface{} // updated by
	UpdatedAt          *gtime.Time // updated at
	DeletedBy          interface{} // deleted by
	DeletedAt          *gtime.Time // deleted at
}
