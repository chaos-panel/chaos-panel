// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-11 20:17:21
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TSequenceDao is the data access object for the table t_sequence.
type TSequenceDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns TSequenceColumns // columns contains all the column names of Table for convenient usage.
}

// TSequenceColumns defines and stores column names for the table t_sequence.
type TSequenceColumns struct {
	Id       string // ID
	Type     string // type
	Sequence string // sequence
}

// tSequenceColumns holds the columns for the table t_sequence.
var tSequenceColumns = TSequenceColumns{
	Id:       "id",
	Type:     "type",
	Sequence: "sequence",
}

// NewTSequenceDao creates and returns a new DAO object for table data access.
func NewTSequenceDao() *TSequenceDao {
	return &TSequenceDao{
		group:   "default",
		table:   "t_sequence",
		columns: tSequenceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TSequenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TSequenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TSequenceDao) Columns() TSequenceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TSequenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TSequenceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TSequenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
