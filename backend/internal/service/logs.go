// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/logs/v1"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
)

type (
	ILogs interface {
		Create(ctx context.Context, log *do.Logs) error
		GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
		GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	}
)

var (
	localLogs ILogs
)

func Logs() ILogs {
	if localLogs == nil {
		panic("implement not found for interface ILogs, forgot register?")
	}
	return localLogs
}

func RegisterLogs(i ILogs) {
	localLogs = i
}
