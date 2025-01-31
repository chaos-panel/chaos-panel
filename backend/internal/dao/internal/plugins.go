// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PluginsDao is the data access object for the table chaosplus_plugins.
type PluginsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns PluginsColumns // columns contains all the column names of Table for convenient usage.
}

// PluginsColumns defines and stores column names for the table chaosplus_plugins.
type PluginsColumns struct {
	Id          string // ID
	TenantId    string // tenant id
	RepoId      string // repo id
	Kind        string // kind
	Type        string // type
	Name        string // name
	Author      string // author
	VersionCode string // version code
	VersionName string // version name
	Desc        string // desc
	CreatedBy   string // created by
	CreatedAt   string // created at
	UpdatedBy   string // updated by
	UpdatedAt   string // updated at
	DeletedBy   string // deleted by
	DeletedAt   string // deleted at
}

// pluginsColumns holds the columns for the table chaosplus_plugins.
var pluginsColumns = PluginsColumns{
	Id:          "id",
	TenantId:    "tenant_id",
	RepoId:      "repo_id",
	Kind:        "kind",
	Type:        "type",
	Name:        "name",
	Author:      "author",
	VersionCode: "version_code",
	VersionName: "version_name",
	Desc:        "desc",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewPluginsDao creates and returns a new DAO object for table data access.
func NewPluginsDao() *PluginsDao {
	return &PluginsDao{
		group:   "default",
		table:   "chaosplus_plugins",
		columns: pluginsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PluginsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PluginsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PluginsDao) Columns() PluginsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PluginsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PluginsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PluginsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
