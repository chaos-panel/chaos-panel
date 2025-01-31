// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResourcesDao is the data access object for the table chaosplus_resources.
type ResourcesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns ResourcesColumns // columns contains all the column names of Table for convenient usage.
}

// ResourcesColumns defines and stores column names for the table chaosplus_resources.
type ResourcesColumns struct {
	Id            string // ID
	TenantId      string // tenant id
	Path          string // path
	PathDepth     string // path depth
	PathMd5       string // path md5
	PathSha1      string // path sha1
	PathSha256    string // path sha256
	PathSha512    string // path sha512
	ChildrenCount string // children count
	CreatedBy     string // created by
	CreatedAt     string // created at
	UpdatedBy     string // updated by
	UpdatedAt     string // updated at
	DeletedBy     string // deleted by
	DeletedAt     string // deleted at
}

// resourcesColumns holds the columns for the table chaosplus_resources.
var resourcesColumns = ResourcesColumns{
	Id:            "id",
	TenantId:      "tenant_id",
	Path:          "path",
	PathDepth:     "path_depth",
	PathMd5:       "path_md5",
	PathSha1:      "path_sha1",
	PathSha256:    "path_sha256",
	PathSha512:    "path_sha512",
	ChildrenCount: "children_count",
	CreatedBy:     "created_by",
	CreatedAt:     "created_at",
	UpdatedBy:     "updated_by",
	UpdatedAt:     "updated_at",
	DeletedBy:     "deleted_by",
	DeletedAt:     "deleted_at",
}

// NewResourcesDao creates and returns a new DAO object for table data access.
func NewResourcesDao() *ResourcesDao {
	return &ResourcesDao{
		group:   "default",
		table:   "chaosplus_resources",
		columns: resourcesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ResourcesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ResourcesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ResourcesDao) Columns() ResourcesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ResourcesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ResourcesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ResourcesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
