package dlock

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
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
		_, err := dao.DistributedLock.Ctx(ctx).Where("created_at <= ?", gtime.Now().Add(-ttl)).Delete()
		if err != nil {
			return err
		}
		var exists *do.DistributedLock
		err = dao.DistributedLock.Ctx(ctx).Where("lock_key = ?", key).Scan(&exists)
		if err != nil {
			return err
		}
		if exists == nil {
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
			// if err is not duplicate entry, return err
			if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
				return err
			}
		}
		// get the exists lock
		exists = nil
		err = dao.DistributedLock.Ctx(ctx).Where("lock_key = ?", key).Scan(&exists)
		if err != nil {
			return err
		}
		if exists == nil {
			return errors.New("add or get lock failed")
		}
		holder := gvar.New(exists.CreatedBy).String()
		// is the exists lock by is not equal to by, return err
		if holder != by {
			return errors.New("lock is already holded by: " + holder + ", not " + by)
		}
		// update the exists lock ttl
		_, err = dao.DistributedLock.Ctx(ctx).Where("lock_key = ?", key).Update(&do.DistributedLock{
			LockKey:   key,
			HostInfo:  host,
			ExpiredAt: gtime.Now().Add(ttl),
			UpdatedBy: by,
			UpdatedAt: gtime.Now(),
		})
		return err
	})
}

func (s *sDlock) Unlock(ctx context.Context, key string, by string) error {
	return dao.DistributedLock.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.DistributedLock.Ctx(ctx).Where("lock_key = ? and created_by = ?", key, by).Delete()
		return err
	})
}
