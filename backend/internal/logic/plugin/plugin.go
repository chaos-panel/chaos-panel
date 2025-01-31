package plugin

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/plugins/v1"
	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/database/gdb"
)

// Plugin 定义插件接口
type Plugin interface {
}

type sPlugin struct{}

func init() {
	service.RegisterPlugin(&sPlugin{})
}

func New() service.IPlugin {
	return &sPlugin{}
}

func (s *sPlugin) Create(ctx context.Context, log *do.Plugins) error {
	log.Id = service.Guid().NextId()
	return dao.Logs.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Model(&entity.Logs{}).Insert(log)
		return err
	})
}

func (s *sPlugin) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Logs.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&res)
	return
}

func (s *sPlugin) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var entities []entity.Plugins
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
