// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileStoragesDao is the data access object for the table chaosplus_file_storages.
type FileStoragesDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns FileStoragesColumns // columns contains all the column names of Table for convenient usage.
}

// FileStoragesColumns defines and stores column names for the table chaosplus_file_storages.
type FileStoragesColumns struct {
	Id            string // ID
	ContentLength string // content length
	ContentMd5    string // content md5
	ContentSha1   string // content sha1
	ContentSha256 string // content sha256
	ContentSha512 string // content sha512
	StorageUri    string // storage uri
	CreatedBy     string // created by
	CreatedAt     string // created at
	UpdatedBy     string // updated by
	UpdatedAt     string // updated at
	DeletedBy     string // deleted by
	DeletedAt     string // deleted at
}

// fileStoragesColumns holds the columns for the table chaosplus_file_storages.
var fileStoragesColumns = FileStoragesColumns{
	Id:            "id",
	ContentLength: "content_length",
	ContentMd5:    "content_md5",
	ContentSha1:   "content_sha1",
	ContentSha256: "content_sha256",
	ContentSha512: "content_sha512",
	StorageUri:    "storage_uri",
	CreatedBy:     "created_by",
	CreatedAt:     "created_at",
	UpdatedBy:     "updated_by",
	UpdatedAt:     "updated_at",
	DeletedBy:     "deleted_by",
	DeletedAt:     "deleted_at",
}

// NewFileStoragesDao creates and returns a new DAO object for table data access.
func NewFileStoragesDao() *FileStoragesDao {
	return &FileStoragesDao{
		group:   "default",
		table:   "chaosplus_file_storages",
		columns: fileStoragesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FileStoragesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FileStoragesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FileStoragesDao) Columns() FileStoragesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FileStoragesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FileStoragesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FileStoragesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
