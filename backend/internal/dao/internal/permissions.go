// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionsDao is the data access object for the table chaosplus_permissions.
type PermissionsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PermissionsColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionsColumns defines and stores column names for the table chaosplus_permissions.
type PermissionsColumns struct {
	Id         string // ID
	RoleId     string // role id
	ResourceId string // resource id
	Action     string // action
	Effect     string // effect
	CreatedBy  string // created by
	CreatedAt  string // created at
	UpdatedBy  string // updated by
	UpdatedAt  string // updated at
}

// permissionsColumns holds the columns for the table chaosplus_permissions.
var permissionsColumns = PermissionsColumns{
	Id:         "id",
	RoleId:     "role_id",
	ResourceId: "resource_id",
	Action:     "action",
	Effect:     "effect",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

// NewPermissionsDao creates and returns a new DAO object for table data access.
func NewPermissionsDao() *PermissionsDao {
	return &PermissionsDao{
		group:   "default",
		table:   "chaosplus_permissions",
		columns: permissionsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PermissionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PermissionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PermissionsDao) Columns() PermissionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PermissionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PermissionsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PermissionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
