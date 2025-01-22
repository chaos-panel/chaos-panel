// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WorkerIdsDao is the data access object for the table chaosplus_worker_ids.
type WorkerIdsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns WorkerIdsColumns // columns contains all the column names of Table for convenient usage.
}

// WorkerIdsColumns defines and stores column names for the table chaosplus_worker_ids.
type WorkerIdsColumns struct {
	Id        string // id
	HostInfo  string // host info
	ExpiredAt string // expired at
	CreatedBy string // created by
	CreatedAt string // locked at
	UpdatedBy string // updated by
	UpdatedAt string // time updated
}

// workerIdsColumns holds the columns for the table chaosplus_worker_ids.
var workerIdsColumns = WorkerIdsColumns{
	Id:        "id",
	HostInfo:  "host_info",
	ExpiredAt: "expired_at",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewWorkerIdsDao creates and returns a new DAO object for table data access.
func NewWorkerIdsDao() *WorkerIdsDao {
	return &WorkerIdsDao{
		group:   "default",
		table:   "chaosplus_worker_ids",
		columns: workerIdsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WorkerIdsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WorkerIdsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WorkerIdsDao) Columns() WorkerIdsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WorkerIdsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WorkerIdsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WorkerIdsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
