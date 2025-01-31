package dfs

import (
	"context"
	"time"

	v1 "github.com/chaos-plus/chaos-plus/api/files/v1"
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

type sDfs struct{}

func init() {
	service.RegisterDfs(&sDfs{})
}

func New() service.IDfs {
	return &sDfs{}
}

func (s *sDfs) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	do, err := converter.ToAny[do.Terminals](*req)
	if err != nil {
		return
	}
	r := g.RequestFromCtx(ctx)
	uid := middleware.GetUserId(r)
	do.Id = service.Guid().NextId()
	do.TenantId = 0
	do.CreatedBy = uid
	do.UpdatedBy = uid
	do.CreatedAt = gtime.Now()
	do.UpdatedAt = gtime.Now()
	do.DeletedBy = 0
	do.DeletedAt = gtime.NewFromTime(time.Unix(0, 0))
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

func (s *sDfs) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Terminals.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&res)
	return
}

func (s *sDfs) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
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

func (s *sDfs) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
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
