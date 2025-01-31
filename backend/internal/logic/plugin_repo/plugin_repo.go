package pluginrepo

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/plugin_repos/v1"
	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/database/gdb"
)

type sPluginRepo struct{}

func init() {
	service.RegisterPluginRepo(&sPluginRepo{})
}

func New() service.IPluginRepo {
	return &sPluginRepo{}
}

func (s *sPluginRepo) Create(ctx context.Context, log *do.PluginRepos) error {
	log.Id = service.Guid().NextId()
	return dao.Logs.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Model(&entity.Logs{}).Insert(log)
		return err
	})
}

func (s *sPluginRepo) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Logs.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&res)
	return
}

func (s *sPluginRepo) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var entities []entity.PluginRepos
	var count int
	m := dao.Logs.Ctx(ctx)
	m = page.ParseQueriesFromCtx(ctx, m, req)
	// m = page.NonDeleted(m)
	// m = page.NonLocked(m)
	err = m.ScanAndCount(&entities, &count, true)
	res = &v1.GetListRes{}
	res.Offset = req.Offset
	res.Limit = req.Limit
	res.Count = len(entities)
	res.Total = count
	res.List = entities
	return
}
