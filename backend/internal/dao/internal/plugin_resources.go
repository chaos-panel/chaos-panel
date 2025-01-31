// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PluginResourcesDao is the data access object for the table chaosplus_plugin_resources.
type PluginResourcesDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns PluginResourcesColumns // columns contains all the column names of Table for convenient usage.
}

// PluginResourcesColumns defines and stores column names for the table chaosplus_plugin_resources.
type PluginResourcesColumns struct {
	Id          string // ID
	PluginId    string // plugin id
	FileType    string // file type
	FileName    string // file name
	FileId      string // file id
	MountType   string // mount type
	MountPath   string // mount path
	MountPolicy string // mount policy
	CreatedBy   string // created by
	CreatedAt   string // created at
	UpdatedBy   string // updated by
	UpdatedAt   string // updated at
	DeletedBy   string // deleted by
	DeletedAt   string // deleted at
}

// pluginResourcesColumns holds the columns for the table chaosplus_plugin_resources.
var pluginResourcesColumns = PluginResourcesColumns{
	Id:          "id",
	PluginId:    "plugin_id",
	FileType:    "file_type",
	FileName:    "file_name",
	FileId:      "file_id",
	MountType:   "mount_type",
	MountPath:   "mount_path",
	MountPolicy: "mount_policy",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewPluginResourcesDao creates and returns a new DAO object for table data access.
func NewPluginResourcesDao() *PluginResourcesDao {
	return &PluginResourcesDao{
		group:   "default",
		table:   "chaosplus_plugin_resources",
		columns: pluginResourcesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PluginResourcesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PluginResourcesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PluginResourcesDao) Columns() PluginResourcesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PluginResourcesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PluginResourcesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PluginResourcesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
