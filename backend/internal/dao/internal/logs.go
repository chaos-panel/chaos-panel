// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogsDao is the data access object for the table chaosplus_logs.
type LogsDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns LogsColumns // columns contains all the column names of Table for convenient usage.
}

// LogsColumns defines and stores column names for the table chaosplus_logs.
type LogsColumns struct {
	Id        string // ID
	Status    string // 状态
	Log       string // 日志
	CreatedBy string // 创建者
	CreatedAt string // 创建时间
}

// logsColumns holds the columns for the table chaosplus_logs.
var logsColumns = LogsColumns{
	Id:        "id",
	Status:    "status",
	Log:       "log",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
}

// NewLogsDao creates and returns a new DAO object for table data access.
func NewLogsDao() *LogsDao {
	return &LogsDao{
		group:   "default",
		table:   "chaosplus_logs",
		columns: logsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LogsDao) Columns() LogsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *LogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
