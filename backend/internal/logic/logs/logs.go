package logs

import (
	"context"

	"github.com/chaos-panel/chaos-panel/internal/dao"
	"github.com/chaos-panel/chaos-panel/internal/model/do"
	"github.com/chaos-panel/chaos-panel/internal/service"
	"github.com/chaos-panel/chaos-panel/utility/guid"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gres"

	_ "github.com/chaos-panel/chaos-panel/internal/packed"
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
