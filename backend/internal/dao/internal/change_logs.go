// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-14 17:54:17
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChangeLogsDao is the data access object for the table chaosplus_change_logs.
type ChangeLogsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns ChangeLogsColumns // columns contains all the column names of Table for convenient usage.
}

// ChangeLogsColumns defines and stores column names for the table chaosplus_change_logs.
type ChangeLogsColumns struct {
	Id        string // id
	UnionId   string // union id
	Snapshot  string // snapshot
	CreatedBy string // created by
	CreatedAt string // locked at
	DeletedBy string // deleted by
	DeletedAt string // time deleted
}

// changeLogsColumns holds the columns for the table chaosplus_change_logs.
var changeLogsColumns = ChangeLogsColumns{
	Id:        "id",
	UnionId:   "union_id",
	Snapshot:  "snapshot",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewChangeLogsDao creates and returns a new DAO object for table data access.
func NewChangeLogsDao() *ChangeLogsDao {
	return &ChangeLogsDao{
		group:   "default",
		table:   "chaosplus_change_logs",
		columns: changeLogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ChangeLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ChangeLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ChangeLogsDao) Columns() ChangeLogsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ChangeLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ChangeLogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ChangeLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
