// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResourcesClosureDao is the data access object for the table chaosplus_resources_closure.
type ResourcesClosureDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns ResourcesClosureColumns // columns contains all the column names of Table for convenient usage.
}

// ResourcesClosureColumns defines and stores column names for the table chaosplus_resources_closure.
type ResourcesClosureColumns struct {
	AncestorId   string // ancestor id
	DescendantId string // descendant id
	Distance     string // distance
}

// resourcesClosureColumns holds the columns for the table chaosplus_resources_closure.
var resourcesClosureColumns = ResourcesClosureColumns{
	AncestorId:   "ancestor_id",
	DescendantId: "descendant_id",
	Distance:     "distance",
}

// NewResourcesClosureDao creates and returns a new DAO object for table data access.
func NewResourcesClosureDao() *ResourcesClosureDao {
	return &ResourcesClosureDao{
		group:   "default",
		table:   "chaosplus_resources_closure",
		columns: resourcesClosureColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ResourcesClosureDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ResourcesClosureDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ResourcesClosureDao) Columns() ResourcesClosureColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ResourcesClosureDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ResourcesClosureDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ResourcesClosureDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
