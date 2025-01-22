// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileStorages is the golang structure of table chaosplus_file_storages for DAO operations like Where/Data.
type FileStorages struct {
	g.Meta        `orm:"table:chaosplus_file_storages, do:true"`
	Id            interface{} // ID
	ContentLength interface{} // content length
	ContentMd5    interface{} // content md5
	ContentSha1   interface{} // content sha1
	ContentSha256 interface{} // content sha256
	ContentSha512 interface{} // content sha512
	StorageUri    interface{} // storage uri
	CreatedBy     interface{} // created by
	CreatedAt     *gtime.Time // created at
	UpdatedBy     interface{} // updated by
	UpdatedAt     *gtime.Time // updated at
	DeletedBy     interface{} // deleted by
	DeletedAt     *gtime.Time // deleted at
}
