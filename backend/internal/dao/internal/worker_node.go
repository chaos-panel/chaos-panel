// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WorkerNodeDao is the data access object for the table chaosplus_worker_node.
type WorkerNodeDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns WorkerNodeColumns // columns contains all the column names of Table for convenient usage.
}

// WorkerNodeColumns defines and stores column names for the table chaosplus_worker_node.
type WorkerNodeColumns struct {
	Id          string // 节点id
	Port        string // 端口
	OsName      string // 系统名称
	OsArch      string // 系统平台
	OsVersion   string // 系统版本
	HostName    string // 主机名称
	Docker      string // IP地址
	Ip          string // IP地址
	Mac         string // MAC地址
	Tag         string // 节点标签
	Desc        string // 节点描述
	TimeCreated string // 创建时间
	TimeUpdated string // 更新时间
}

// workerNodeColumns holds the columns for the table chaosplus_worker_node.
var workerNodeColumns = WorkerNodeColumns{
	Id:          "id",
	Port:        "port",
	OsName:      "os_name",
	OsArch:      "os_arch",
	OsVersion:   "os_version",
	HostName:    "host_name",
	Docker:      "docker",
	Ip:          "ip",
	Mac:         "mac",
	Tag:         "tag",
	Desc:        "desc",
	TimeCreated: "time_created",
	TimeUpdated: "time_updated",
}

// NewWorkerNodeDao creates and returns a new DAO object for table data access.
func NewWorkerNodeDao() *WorkerNodeDao {
	return &WorkerNodeDao{
		group:   "default",
		table:   "chaosplus_worker_node",
		columns: workerNodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WorkerNodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WorkerNodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WorkerNodeDao) Columns() WorkerNodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WorkerNodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WorkerNodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WorkerNodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
