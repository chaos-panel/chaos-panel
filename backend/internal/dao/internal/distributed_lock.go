// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-07 10:37:30
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DistributedLockDao is the data access object for the table chaosplus_distributed_lock.
type DistributedLockDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns DistributedLockColumns // columns contains all the column names of Table for convenient usage.
}

// DistributedLockColumns defines and stores column names for the table chaosplus_distributed_lock.
type DistributedLockColumns struct {
	LockKey   string // lock key
	HostInfo  string // host info
	ExpiredAt string // expired at
	CreatedBy string // created by
	CreatedAt string // locked at
	UpdatedBy string // updated by
	UpdatedAt string // time updated
}

// distributedLockColumns holds the columns for the table chaosplus_distributed_lock.
var distributedLockColumns = DistributedLockColumns{
	LockKey:   "lock_key",
	HostInfo:  "host_info",
	ExpiredAt: "expired_at",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewDistributedLockDao creates and returns a new DAO object for table data access.
func NewDistributedLockDao() *DistributedLockDao {
	return &DistributedLockDao{
		group:   "default",
		table:   "chaosplus_distributed_lock",
		columns: distributedLockColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DistributedLockDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DistributedLockDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DistributedLockDao) Columns() DistributedLockColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DistributedLockDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DistributedLockDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DistributedLockDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
