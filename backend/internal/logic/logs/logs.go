package logs

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/logs/v1"
	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/database/gdb"
)

type sLogs struct{}

func init() {
	service.RegisterLogs(&sLogs{})
}

func New() service.ILogs {
	return &sLogs{}
}

func (s *sLogs) Create(ctx context.Context, log *do.Logs) error {
	log.Id = service.Guid().NextId()
	return dao.Logs.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Model(&entity.Logs{}).Insert(log)
		return err
	})
}

func (s *sLogs) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Logs.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&res)
	return
}

func (s *sLogs) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var entities []entity.Logs
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
