package logs

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/logs/v1"
	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/guid"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/database/gdb"

	_ "github.com/chaos-plus/chaos-plus/internal/packed"
)

type sLogs struct{}

func init() {
	service.RegisterLogs(&sLogs{})
}

func New() service.ILogs {
	return &sLogs{}
}

func (s *sLogs) Create(ctx context.Context, log *do.Logs) error {
	log.Id = guid.NextId()
	dao.Logs.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Logs.Ctx(ctx).Insert(log)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (s *sLogs) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var entities []entity.Logs
	var count int
	m := dao.Logs.Ctx(ctx)
	m = page.ParseQueriesFromCtx(ctx, m, req)
	// m = page.NonDeleted(m)
	// m = page.NonLocked(m)
	m.ScanAndCount(&entities, &count, true)
	res = &v1.GetListRes{}
	res.Offset = req.Offset
	res.Limit = req.Limit
	res.Count = len(entities)
	res.Total = count
	res.List = entities
	return
}

func (s *sLogs) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Logs.Ctx(ctx)
	m.Where("id = ?", req.Id).Scan(&res)
	return
}
