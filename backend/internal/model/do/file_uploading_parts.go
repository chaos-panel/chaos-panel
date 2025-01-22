// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileUploadingParts is the golang structure of table chaosplus_file_uploading_parts for DAO operations like Where/Data.
type FileUploadingParts struct {
	g.Meta    `orm:"table:chaosplus_file_uploading_parts, do:true"`
	Id        interface{} // ID
	ExpiredAt *gtime.Time // expired at
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // updated at
}
