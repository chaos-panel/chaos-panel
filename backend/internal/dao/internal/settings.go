// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-18 00:02:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingsDao is the data access object for the table chaosplus_settings.
type SettingsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SettingsColumns // columns contains all the column names of Table for convenient usage.
}

// SettingsColumns defines and stores column names for the table chaosplus_settings.
type SettingsColumns struct {
	Id        string // id
	TenantId  string // tenant id
	Group     string // group
	Key       string // key
	KeyName   string // key display
	Val       string // value
	ValType   string // value type
	CreatedBy string // created by
	CreatedAt string // locked at
	UpdatedBy string // updated by
	UpdatedAt string // time updated
	DeletedBy string // deleted by
	DeletedAt string // time deleted
}

// settingsColumns holds the columns for the table chaosplus_settings.
var settingsColumns = SettingsColumns{
	Id:        "id",
	TenantId:  "tenant_id",
	Group:     "group",
	Key:       "key",
	KeyName:   "key_name",
	Val:       "val",
	ValType:   "val_type",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewSettingsDao creates and returns a new DAO object for table data access.
func NewSettingsDao() *SettingsDao {
	return &SettingsDao{
		group:   "default",
		table:   "chaosplus_settings",
		columns: settingsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SettingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SettingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SettingsDao) Columns() SettingsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SettingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SettingsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SettingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
