package dlock

import (
	"context"
	"strings"
	"time"

	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDlock struct{}

func init() {
	service.RegisterDlock(&sDlock{})
}

func New() service.IDlock {
	return &sDlock{}
}

func (s *sDlock) Lock(ctx context.Context, key string, ttl time.Duration, by string, host string) error {
	return dao.DistributedLock.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.DistributedLock.Ctx(ctx).Where("created_at <= DATE_SUB(NOW(), INTERVAL ? SECOND)", ttl).Delete()
		if err != nil {
			return err
		}
		// try insert lock
		_, err = dao.DistributedLock.Ctx(ctx).Insert(&do.DistributedLock{
			LockKey:   key,
			HostInfo:  host,
			ExpiredAt: gtime.Now().Add(ttl),
			CreatedBy: by,
			CreatedAt: gtime.Now(),
			UpdatedBy: by,
			UpdatedAt: gtime.Now(),
		})
		if err == nil {
			return nil
		}
		// if err is not duplicate entry, return err
		if !strings.Contains(err.Error(), "Duplicate entry") {
			return err
		}
		// get the exists lock
		exists := do.DistributedLock{}
		scanserr := dao.DistributedLock.Ctx(ctx).Where("lock_key = ?", key).Scan(&exists)
		if scanserr != nil {
			return scanserr
		}
		// is the exists lock by is not equal to by, return err
		if exists.CreatedBy != by {
			g.Log().Error(ctx, "33333----lock error: ", exists.CreatedBy)
			g.Log().Error(ctx, "44444----lock error: ", by)
			g.Log().Error(ctx, "55555----lock error: ", exists.HostInfo)
			g.Log().Error(ctx, "66666----lock error: ", host)
			return err
		}
		// update the exists lock ttl
		_, updateerr := dao.DistributedLock.Ctx(ctx).Where("lock_key = ?", key).Update(&do.DistributedLock{
			LockKey:   key,
			HostInfo:  host,
			ExpiredAt: gtime.Now().Add(ttl),
			UpdatedBy: by,
			UpdatedAt: gtime.Now(),
		})
		if updateerr == nil {
			return nil
		}
		return updateerr
	})
}

func (s *sDlock) Unlock(ctx context.Context, key string, by string) error {
	return dao.DistributedLock.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.DistributedLock.Ctx(ctx).Where("lock_key = ? and created_by = ?", key, by).Delete()
		return err
	})
}
