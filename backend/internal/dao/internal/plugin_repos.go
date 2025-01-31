// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PluginReposDao is the data access object for the table chaosplus_plugin_repos.
type PluginReposDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PluginReposColumns // columns contains all the column names of Table for convenient usage.
}

// PluginReposColumns defines and stores column names for the table chaosplus_plugin_repos.
type PluginReposColumns struct {
	Id        string // ID
	TenantId  string // tenant id
	Name      string // name
	Url       string // url
	Desc      string // desc
	CreatedBy string // created by
	CreatedAt string // created at
	UpdatedBy string // updated by
	UpdatedAt string // updated at
	DeletedBy string // deleted by
	DeletedAt string // deleted at
}

// pluginReposColumns holds the columns for the table chaosplus_plugin_repos.
var pluginReposColumns = PluginReposColumns{
	Id:        "id",
	TenantId:  "tenant_id",
	Name:      "name",
	Url:       "url",
	Desc:      "desc",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	DeletedBy: "deleted_by",
	DeletedAt: "deleted_at",
}

// NewPluginReposDao creates and returns a new DAO object for table data access.
func NewPluginReposDao() *PluginReposDao {
	return &PluginReposDao{
		group:   "default",
		table:   "chaosplus_plugin_repos",
		columns: pluginReposColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PluginReposDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PluginReposDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PluginReposDao) Columns() PluginReposColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PluginReposDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PluginReposDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PluginReposDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
