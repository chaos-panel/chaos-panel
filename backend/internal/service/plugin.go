// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/plugins/v1"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
)

type (
	IPlugin interface {
		Create(ctx context.Context, log *do.Plugins) error
		GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
		GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	}
)

var (
	localPlugin IPlugin
)

func Plugin() IPlugin {
	if localPlugin == nil {
		panic("implement not found for interface IPlugin, forgot register?")
	}
	return localPlugin
}

func RegisterPlugin(i IPlugin) {
	localPlugin = i
}
