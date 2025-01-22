// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TerminalsDao is the data access object for the table chaosplus_terminals.
type TerminalsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns TerminalsColumns // columns contains all the column names of Table for convenient usage.
}

// TerminalsColumns defines and stores column names for the table chaosplus_terminals.
type TerminalsColumns struct {
	Id        string // ID
	TenantId  string // tenant id
	Type      string // type
	Host      string // host
	Port      string // port
	Username  string // username
	Password  string // password
	CreatedBy string // created by
	CreatedAt string // created at
	UpdatedBy string // updated by
	UpdatedAt string // updated at
	DeletedBy string // deleted by
	DeletedAt string // deleted at
}

// terminalsColumns holds the columns for the table chaosplus_terminals.
var terminalsColumns = TerminalsColumns{
	Id:        "id",
	TenantId:  "tenant_id",
	Type:      "type",
	Host:      "host",
	Port:      "port",
	Username:  "username",
	Password:  "password",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewTerminalsDao creates and returns a new DAO object for table data access.
func NewTerminalsDao() *TerminalsDao {
	return &TerminalsDao{
		group:   "default",
		table:   "chaosplus_terminals",
		columns: terminalsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TerminalsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TerminalsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TerminalsDao) Columns() TerminalsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TerminalsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TerminalsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TerminalsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
