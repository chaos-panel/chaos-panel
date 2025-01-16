// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DistributedLocksDao is the data access object for the table chaosplus_distributed_locks.
type DistributedLocksDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns DistributedLocksColumns // columns contains all the column names of Table for convenient usage.
}

// DistributedLocksColumns defines and stores column names for the table chaosplus_distributed_locks.
type DistributedLocksColumns struct {
	LockKey   string // lock key
	HostInfo  string // host info
	ExpiredAt string // expired at
	CreatedBy string // created by
	CreatedAt string // locked at
	UpdatedBy string // updated by
	UpdatedAt string // time updated
}

// distributedLocksColumns holds the columns for the table chaosplus_distributed_locks.
var distributedLocksColumns = DistributedLocksColumns{
	LockKey:   "lock_key",
	HostInfo:  "host_info",
	ExpiredAt: "expired_at",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewDistributedLocksDao creates and returns a new DAO object for table data access.
func NewDistributedLocksDao() *DistributedLocksDao {
	return &DistributedLocksDao{
		group:   "default",
		table:   "chaosplus_distributed_locks",
		columns: distributedLocksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DistributedLocksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DistributedLocksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DistributedLocksDao) Columns() DistributedLocksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DistributedLocksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DistributedLocksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DistributedLocksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
