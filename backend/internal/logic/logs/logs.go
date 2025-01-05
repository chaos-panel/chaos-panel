package logs

import (
	"context"

	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/guid"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gres"

	_ "github.com/chaos-plus/chaos-plus/internal/packed"
)

type sLogs struct{}

func init() {
	service.RegisterLogs(&sLogs{})
	gres.Dump()
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
