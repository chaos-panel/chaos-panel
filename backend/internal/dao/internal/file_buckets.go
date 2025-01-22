// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileBucketsDao is the data access object for the table chaosplus_file_buckets.
type FileBucketsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns FileBucketsColumns // columns contains all the column names of Table for convenient usage.
}

// FileBucketsColumns defines and stores column names for the table chaosplus_file_buckets.
type FileBucketsColumns struct {
	Id        string // ID
	TenantId  string // tenant id
	Name      string // name
	Capacity  string // capacity
	CreatedBy string // created by
	CreatedAt string // created at
	UpdatedBy string // updated by
	UpdatedAt string // updated at
	DeletedBy string // deleted by
	DeletedAt string // deleted at
}

// fileBucketsColumns holds the columns for the table chaosplus_file_buckets.
var fileBucketsColumns = FileBucketsColumns{
	Id:        "id",
	TenantId:  "tenant_id",
	Name:      "name",
	Capacity:  "capacity",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewFileBucketsDao creates and returns a new DAO object for table data access.
func NewFileBucketsDao() *FileBucketsDao {
	return &FileBucketsDao{
		group:   "default",
		table:   "chaosplus_file_buckets",
		columns: fileBucketsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FileBucketsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FileBucketsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FileBucketsDao) Columns() FileBucketsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FileBucketsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FileBucketsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FileBucketsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
