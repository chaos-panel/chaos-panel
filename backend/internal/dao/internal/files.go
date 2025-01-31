// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FilesDao is the data access object for the table chaosplus_files.
type FilesDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns FilesColumns // columns contains all the column names of Table for convenient usage.
}

// FilesColumns defines and stores column names for the table chaosplus_files.
type FilesColumns struct {
	Id                 string // ID
	BucketId           string // bucket id
	IsDir              string // is dir
	IsSystem           string // is system file
	IsHidden           string // is hidden file
	Size               string // size
	Name               string // name
	Path               string // path
	PathDepth          string // path depth
	PathMd5            string // path md5
	PathSha1           string // path sha1
	PathSha256         string // path sha256
	PathSha512         string // path sha512
	Ext                string // extension name
	Type               string // type
	Subtype            string // subtype
	FileId             string // physical file id
	Desc               string // description
	SubFileCount       string // sub file node count
	SubFileSystemCount string // sub file system node count
	SubFileHiddenCount string // sub file hidden node count
	SubDirCount        string // sub dir node count
	SubDirSystemCount  string // sub dir system node count
	SubDirHiddenCount  string // sub dir hidden node count
	CreatedBy          string // created by
	CreatedAt          string // created at
	UpdatedBy          string // updated by
	UpdatedAt          string // updated at
	DeletedBy          string // deleted by
	DeletedAt          string // deleted at
}

// filesColumns holds the columns for the table chaosplus_files.
var filesColumns = FilesColumns{
	Id:                 "id",
	BucketId:           "bucket_id",
	IsDir:              "is_dir",
	IsSystem:           "is_system",
	IsHidden:           "is_hidden",
	Size:               "size",
	Name:               "name",
	Path:               "path",
	PathDepth:          "path_depth",
	PathMd5:            "path_md5",
	PathSha1:           "path_sha1",
	PathSha256:         "path_sha256",
	PathSha512:         "path_sha512",
	Ext:                "ext",
	Type:               "type",
	Subtype:            "subtype",
	FileId:             "file_id",
	Desc:               "desc",
	SubFileCount:       "sub_file_count",
	SubFileSystemCount: "sub_file_system_count",
	SubFileHiddenCount: "sub_file_hidden_count",
	SubDirCount:        "sub_dir_count",
	SubDirSystemCount:  "sub_dir_system_count",
	SubDirHiddenCount:  "sub_dir_hidden_count",
	CreatedBy:          "created_by",
	CreatedAt:          "created_at",
	UpdatedBy:          "updated_by",
	UpdatedAt:          "updated_at",
	DeletedBy:          "deleted_by",
	DeletedAt:          "deleted_at",
}

// NewFilesDao creates and returns a new DAO object for table data access.
func NewFilesDao() *FilesDao {
	return &FilesDao{
		group:   "default",
		table:   "chaosplus_files",
		columns: filesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FilesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FilesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FilesDao) Columns() FilesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FilesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FilesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FilesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
