// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthTokensDao is the data access object for the table chaosplus_auth_tokens.
type AuthTokensDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns AuthTokensColumns // columns contains all the column names of Table for convenient usage.
}

// AuthTokensColumns defines and stores column names for the table chaosplus_auth_tokens.
type AuthTokensColumns struct {
	Id                    string // ID
	TenantId              string // tenant id
	UserId                string // user id
	ClientInfo            string // client info
	AccessToken           string // access token
	AccessTokenJti        string // access token jti
	AccessTokenExpiredAt  string // access token expires at
	RefreshToken          string // refresh token
	RefreshTokenExpiredAt string // refresh token expires at
	CreatedBy             string // created by
	CreatedAt             string // created at
	UpdatedBy             string // updated by
	UpdatedAt             string // updated at
	DeletedBy             string // deleted by
	DeletedAt             string // deleted at
}

// authTokensColumns holds the columns for the table chaosplus_auth_tokens.
var authTokensColumns = AuthTokensColumns{
	Id:                    "id",
	TenantId:              "tenant_id",
	UserId:                "user_id",
	ClientInfo:            "client_info",
	AccessToken:           "access_token",
	AccessTokenJti:        "access_token_jti",
	AccessTokenExpiredAt:  "access_token_expired_at",
	RefreshToken:          "refresh_token",
	RefreshTokenExpiredAt: "refresh_token_expired_at",
	CreatedBy:             "created_by",
	CreatedAt:             "created_at",
	UpdatedBy:             "updated_by",
	UpdatedAt:             "updated_at",
	DeletedBy:             "deleted_by",
	DeletedAt:             "deleted_at",
}

// NewAuthTokensDao creates and returns a new DAO object for table data access.
func NewAuthTokensDao() *AuthTokensDao {
	return &AuthTokensDao{
		group:   "default",
		table:   "chaosplus_auth_tokens",
		columns: authTokensColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AuthTokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AuthTokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AuthTokensDao) Columns() AuthTokensColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AuthTokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AuthTokensDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AuthTokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
