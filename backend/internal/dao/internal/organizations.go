// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrganizationsDao is the data access object for the table chaosplus_organizations.
type OrganizationsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns OrganizationsColumns // columns contains all the column names of Table for convenient usage.
}

// OrganizationsColumns defines and stores column names for the table chaosplus_organizations.
type OrganizationsColumns struct {
	Id          string // ID
	TenantId    string // tenant id
	Name        string // name
	Code        string // code
	Logo        string // logo
	Email       string // email
	Location    string // location
	Website     string // website
	Description string // description
	Type        string // type
	CreatedBy   string // created by
	CreatedAt   string // created at
	UpdatedBy   string // updated by
	UpdatedAt   string // updated at
	DeletedBy   string // deleted by
	DeletedAt   string // deleted at
}

// organizationsColumns holds the columns for the table chaosplus_organizations.
var organizationsColumns = OrganizationsColumns{
	Id:          "id",
	TenantId:    "tenant_id",
	Name:        "name",
	Code:        "code",
	Logo:        "logo",
	Email:       "email",
	Location:    "location",
	Website:     "website",
	Description: "description",
	Type:        "type",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewOrganizationsDao creates and returns a new DAO object for table data access.
func NewOrganizationsDao() *OrganizationsDao {
	return &OrganizationsDao{
		group:   "default",
		table:   "chaosplus_organizations",
		columns: organizationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrganizationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrganizationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrganizationsDao) Columns() OrganizationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrganizationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrganizationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrganizationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
