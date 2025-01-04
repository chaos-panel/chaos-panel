// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-17 00:12:10
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogsDao is the data access object for table logs.
type LogsDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns LogsColumns // columns contains all the column names of Table for convenient usage.
}

// LogsColumns defines and stores column names for table logs.
type LogsColumns struct {
	Id        string // ID
	Status    string // 状态
	Log       string // 日志
	CreatedBy string // 创建者
	CreatedAt string // 创建时间
}

// logsColumns holds the columns for table logs.
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
		table:   "logs",
		columns: logsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LogsDao) Columns() LogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
