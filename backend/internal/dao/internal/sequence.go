// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SequenceDao is the data access object for the table chaosplus_sequence.
type SequenceDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SequenceColumns // columns contains all the column names of Table for convenient usage.
}

// SequenceColumns defines and stores column names for the table chaosplus_sequence.
type SequenceColumns struct {
	Id       string // ID
	Type     string // type
	Sequence string // sequence
}

// sequenceColumns holds the columns for the table chaosplus_sequence.
var sequenceColumns = SequenceColumns{
	Id:       "id",
	Type:     "type",
	Sequence: "sequence",
}

// NewSequenceDao creates and returns a new DAO object for table data access.
func NewSequenceDao() *SequenceDao {
	return &SequenceDao{
		group:   "default",
		table:   "chaosplus_sequence",
		columns: sequenceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SequenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SequenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SequenceDao) Columns() SequenceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SequenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SequenceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SequenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
