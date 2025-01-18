package guid

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/terminals/v1"
	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/converter"
	"github.com/chaos-plus/chaos-plus/utility/middleware"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sTerminal struct{}

func init() {
	service.RegisterTerminal(&sTerminal{})
}

func New() service.ITerminal {
	return &sTerminal{}
}

func (s *sTerminal) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	do, err := converter.DeepCopy[do.Terminals](*req)
	if err != nil {
		return
	}
	do.Id = service.Guid().NextId()
	err = dao.Terminals.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Model(&entity.Terminals{}).Insert(do)
		return err
	})
	if err != nil {
		return
	}
	dao.Terminals.Ctx(ctx).Where("id = ?", do.Id).Scan(&res)
	return
}

func (s *sTerminal) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Terminals.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&res)
	return
}

func (s *sTerminal) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var entities []entity.Terminals
	var count int
	m := dao.Terminals.Ctx(ctx)
	m = page.ParseQueriesFromCtx(ctx, m, req)
	m = page.NonDeleted(m)
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

func (s *sTerminal) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	uid := middleware.GetUserId(r)
	err = dao.Terminals.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := tx.Model(&entity.Terminals{}).Where("id = ?", req.Id).Update(&do.Terminals{
			DeletedAt: gtime.Now(),
			DeletedBy: uid,
		})
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		res = (*v1.DeleteRes)(&rows)
		return err
	})
	return
}
