// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CasbinRulesDao is the data access object for the table chaosplus_casbin_rules.
type CasbinRulesDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns CasbinRulesColumns // columns contains all the column names of Table for convenient usage.
}

// CasbinRulesColumns defines and stores column names for the table chaosplus_casbin_rules.
type CasbinRulesColumns struct {
	Id        string // ID
	PType     string //
	V0        string //
	V1        string //
	V2        string //
	V3        string //
	V4        string //
	V5        string //
	CreatedAt string //
}

// casbinRulesColumns holds the columns for the table chaosplus_casbin_rules.
var casbinRulesColumns = CasbinRulesColumns{
	Id:        "id",
	PType:     "p_type",
	V0:        "v0",
	V1:        "v1",
	V2:        "v2",
	V3:        "v3",
	V4:        "v4",
	V5:        "v5",
	CreatedAt: "created_at",
}

// NewCasbinRulesDao creates and returns a new DAO object for table data access.
func NewCasbinRulesDao() *CasbinRulesDao {
	return &CasbinRulesDao{
		group:   "default",
		table:   "chaosplus_casbin_rules",
		columns: casbinRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CasbinRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CasbinRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CasbinRulesDao) Columns() CasbinRulesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CasbinRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CasbinRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CasbinRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
