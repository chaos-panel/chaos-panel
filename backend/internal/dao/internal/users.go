// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for the table chaosplus_users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for the table chaosplus_users.
type UsersColumns struct {
	Id           string // ID
	TenantId     string // tenant id
	Username     string // username
	Firstname    string // firstname
	Lastname     string // lastname
	Nickname     string // nickname
	Password     string // password
	Phone        string // phone
	Email        string // email
	AvatarFileId string // avatar file id
	WxUnionId    string // wx union id
	WxmaOpenId   string // wxma open id
	WxmpOpenId   string // wxmp open id
	GithubUserId string // github user id
	LockedBy     string // locked by
	LockedAt     string // locked at
	CreatedBy    string // created by
	CreatedAt    string // created at
	UpdatedBy    string // updated by
	UpdatedAt    string // updated at
	DeletedBy    string // deleted by
	DeletedAt    string // deleted at
}

// usersColumns holds the columns for the table chaosplus_users.
var usersColumns = UsersColumns{
	Id:           "id",
	TenantId:     "tenant_id",
	Username:     "username",
	Firstname:    "firstname",
	Lastname:     "lastname",
	Nickname:     "nickname",
	Password:     "password",
	Phone:        "phone",
	Email:        "email",
	AvatarFileId: "avatar_file_id",
	WxUnionId:    "wx_union_id",
	WxmaOpenId:   "wxma_open_id",
	WxmpOpenId:   "wxmp_open_id",
	GithubUserId: "github_user_id",
	LockedBy:     "locked_by",
	LockedAt:     "locked_at",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	UpdatedBy:    "updated_by",
	UpdatedAt:    "updated_at",
	DeletedBy:    "deleted_by",
	DeletedAt:    "deleted_at",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "chaosplus_users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
