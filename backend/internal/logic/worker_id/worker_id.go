package workerid

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/crypto"
	netutils "github.com/chaos-plus/chaos-plus/utility/net"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/shirou/gopsutil/v3/host"
)

type WorkerId struct {
	Id  uint
	Tag string
}

type sWorkerId struct{}

var (
	cache *WorkerId
	timer *gtimer.Entry
)

func init() {
	service.RegisterWorkerId(&sWorkerId{})
}

func New() service.IWorkerId {
	return &sWorkerId{}
}

func (s *sWorkerId) InitOrPanic(ctx context.Context, extra ...interface{}) {
	const (
		ttl = time.Minute * 15
	)
	getWorkerIdOrPanic(ctx, "", ttl, extra...)
	g.Log().Debug(ctx, "get worker id ok: "+gvar.New(cache.Id).String())
	if timer != nil {
		return
	}
	timer = gtimer.AddSingleton(ctx, ttl/5, func(ctx context.Context) {
		getWorkerIdOrPanic(ctx, cache.Tag, ttl, extra)
	})
}

func (s *sWorkerId) GetWorkerId(ctx context.Context) uint {
	if cache == nil {
		s.InitOrPanic(ctx)
	}
	if cache == nil {
		panic("get worker id failed")
	}
	return cache.Id
}

func getWorkerIdOrPanic(ctx context.Context, tag string, ttl time.Duration, extra ...interface{}) uint {
	workerId, err := getOrRefreshWorkerId(ctx, tag, ttl, extra...)
	if err != nil {
		panic(err)
	}
	if workerId == nil {
		panic("get worker id failed")
	}
	if cache == nil {
		cache = workerId
	}
	if workerId.Id != cache.Id {
		panic("workid changed")
	}
	return workerId.Id
}

func getOrRefreshWorkerId(ctx context.Context, tag string, ttl time.Duration, extra ...interface{}) (*WorkerId, error) {
	hostInfo, _ := host.InfoWithContext(ctx)

	workerInfo := gjson.New(make(map[string]interface{}))
	workerInfo.Set("hostId", hostInfo.HostID)
	workerInfo.Set("hostname", hostInfo.Hostname)
	workerInfo.Set("os", hostInfo.OS)
	workerInfo.Set("os", hostInfo.OS)
	workerInfo.Set("platform", hostInfo.Platform)
	workerInfo.Set("platformFamily", hostInfo.PlatformFamily)
	workerInfo.Set("platformVersion", hostInfo.PlatformVersion)
	workerInfo.Set("nets", netutils.GetLanAll())
	workerInfo.Set("extra", extra)

	workerHost := workerInfo.String()
	workerTag := tag
	if workerTag == "" {
		workerTag = crypto.Sha256(workerHost)
	}
	workerId := gvar.New(0).Uint()

	err := dao.WorkerId.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		const (
			dlockKey = "__.inner.service.workerid.refresh"
		)
		// get distributed locker
		err = service.Dlock().Lock(ctx, dlockKey, time.Minute*5, workerTag, workerHost)
		if err != nil {
			return err
		}
		// release distributed locker
		defer func() {
			unlockerr := service.Dlock().Unlock(ctx, dlockKey, hostInfo.String())
			if unlockerr != nil {
				err = unlockerr
			}
		}()

		// delete expired workerId
		dao.WorkerId.Ctx(ctx).Where("created_at <= ?", gtime.Now().Add(-ttl)).Delete()

		var exists *do.WorkerId

		// get exists
		err = dao.WorkerId.Ctx(ctx).Where("created_by = ?", workerTag).Scan(&exists)
		if err != nil {
			return err
		}
		if exists == nil {
			// get all workerId
			all, _, err := dao.WorkerId.Ctx(ctx).Fields("id").AllAndCount(false)
			if err != nil {
				return err
			}
			// generate new workerId
			if len(all) > 0 {
				ids := make([]uint, len(all))
				for i, v := range all {
					ids[i] = v["id"].Uint()
				}
				for _, v := range ids {
					if v != workerId {
						break
					}
					workerId++
				}

			}
			// add workerId
			_, err = dao.WorkerId.Ctx(ctx).Insert(&do.WorkerId{
				Id:        workerId,
				HostInfo:  workerHost,
				ExpiredAt: gtime.Now().Add(ttl),
				CreatedBy: workerTag,
				CreatedAt: gtime.Now(),
				UpdatedBy: workerTag,
				UpdatedAt: gtime.Now(),
			})
			// if err is not duplicate entry, return err
			if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
				return err
			}
		}
		// get the exists workerId
		exists = nil
		err = dao.WorkerId.Ctx(ctx).Where("id = ?", workerId).Scan(&exists)
		if err != nil {
			return err
		}
		if exists == nil {
			return errors.New("add or get woker id failed")
		}
		holder := gvar.New(exists.CreatedBy).String()
		if holder != workerTag {
			return errors.New("worker id is already holded by: " + holder + ", not " + workerTag)
		}
		// update the exists worker id ttl
		_, err = dao.WorkerId.Ctx(ctx).Where("id = ?", workerId).Update(&do.WorkerId{
			Id:        workerId,
			HostInfo:  workerHost,
			ExpiredAt: gtime.Now().Add(ttl),
			UpdatedBy: workerTag,
			UpdatedAt: gtime.Now(),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &WorkerId{Id: workerId, Tag: workerTag}, nil
}
