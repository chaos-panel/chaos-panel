// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileUploadingPartsDao is the data access object for the table chaosplus_file_uploading_parts.
type FileUploadingPartsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns FileUploadingPartsColumns // columns contains all the column names of Table for convenient usage.
}

// FileUploadingPartsColumns defines and stores column names for the table chaosplus_file_uploading_parts.
type FileUploadingPartsColumns struct {
	Id        string // ID
	ExpiredAt string // expired at
	CreatedBy string // created by
	CreatedAt string // created at
	UpdatedBy string // updated by
	UpdatedAt string // updated at
}

// fileUploadingPartsColumns holds the columns for the table chaosplus_file_uploading_parts.
var fileUploadingPartsColumns = FileUploadingPartsColumns{
	Id:        "id",
	ExpiredAt: "expired_at",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewFileUploadingPartsDao creates and returns a new DAO object for table data access.
func NewFileUploadingPartsDao() *FileUploadingPartsDao {
	return &FileUploadingPartsDao{
		group:   "default",
		table:   "chaosplus_file_uploading_parts",
		columns: fileUploadingPartsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FileUploadingPartsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FileUploadingPartsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FileUploadingPartsDao) Columns() FileUploadingPartsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FileUploadingPartsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FileUploadingPartsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FileUploadingPartsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
