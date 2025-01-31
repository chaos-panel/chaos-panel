// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/plugin_repos/v1"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
)

type (
	IPluginRepo interface {
		Create(ctx context.Context, log *do.PluginRepos) error
		GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
		GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	}
)

var (
	localPluginRepo IPluginRepo
)

func PluginRepo() IPluginRepo {
	if localPluginRepo == nil {
		panic("implement not found for interface IPluginRepo, forgot register?")
	}
	return localPluginRepo
}

func RegisterPluginRepo(i IPluginRepo) {
	localPluginRepo = i
}
