// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-18 00:02:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LabelsDao is the data access object for the table chaosplus_labels.
type LabelsDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns LabelsColumns // columns contains all the column names of Table for convenient usage.
}

// LabelsColumns defines and stores column names for the table chaosplus_labels.
type LabelsColumns struct {
	LabelId    string // label id
	UnionId    string // union id
	LabelIndex string // label index
	LabelName  string // label name
	LabelColor string // label color
	CreatedBy  string // created by
	CreatedAt  string // locked at
}

// labelsColumns holds the columns for the table chaosplus_labels.
var labelsColumns = LabelsColumns{
	LabelId:    "label_id",
	UnionId:    "union_id",
	LabelIndex: "label_index",
	LabelName:  "label_name",
	LabelColor: "label_color",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
}

// NewLabelsDao creates and returns a new DAO object for table data access.
func NewLabelsDao() *LabelsDao {
	return &LabelsDao{
		group:   "default",
		table:   "chaosplus_labels",
		columns: labelsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LabelsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LabelsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LabelsDao) Columns() LabelsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LabelsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LabelsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *LabelsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
