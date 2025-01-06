package workerid

import (
	"context"
	"time"

	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/crypto"
	netutils "github.com/chaos-plus/chaos-plus/utility/net"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shirou/gopsutil/v3/host"
)

const (
	dlockKey    = "service.worker_id"
	persistTime = time.Minute * 5
)

type WokerIdCache struct {
	id  uint
	tag string
}

type sWorkerId struct{}

var (
	cache = &WokerIdCache{}
)

func init() {
	service.RegisterWorkerId(&sWorkerId{})
}

func New() service.IWorkerId {
	return &sWorkerId{}
}

func (s *sWorkerId) RefreshWorkerId() {
}

func (s *sWorkerId) GetWorkerIdOrPanic(ctx context.Context) uint {
	id, err := s.GetWorkerId(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (s *sWorkerId) GetWorkerId(ctx context.Context) (uint, error) {
	if cache.id > 0 {
		return cache.id, nil
	}

	hostInfo, _ := host.Info()

	workerInfo := gjson.New(make(map[string]interface{}))
	workerInfo.Set("hostId", hostInfo.HostID)
	workerInfo.Set("host", hostInfo)
	workerInfo.Set("host", hostInfo)
	workerInfo.Set("host", hostInfo)
	workerInfo.Set("host", hostInfo)
	workerInfo.Set("host", hostInfo)
	workerInfo.Set("nets", netutils.GetLanAll())

	workerHost := workerInfo.String()
	workerTag := crypto.Sha256(workerInfo.String())

	var workerId uint = 0
	err := dao.WorkerId.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err := service.Dlock().Lock(ctx, dlockKey, time.Minute, workerTag, workerHost)
		if err != nil {
			return err
		}

		// delete expired workerId
		dao.WorkerId.Ctx(ctx).Where("created_at <= ?", gtime.Now().Add(-persistTime)).Delete()
		// get all workerId
		all, _, err := dao.WorkerId.Ctx(ctx).Fields("id").AllAndCount(false)
		if err != nil {
			return err
		}
		g.Log().Print(ctx, "========>workerId count: ", all)
		// generate new workerId

		// add workerId

		// release distributed locker
		err = service.Dlock().Unlock(ctx, dlockKey, hostInfo.String())
		if err != nil {
			return err
		}
		dao.WorkerId.Ctx(ctx).Insert(&do.WorkerId{
			Id:        1,
			HostInfo:  workerHost,
			ExpiredAt: gtime.Now().Add(persistTime),
			CreatedBy: workerTag,
			CreatedAt: gtime.Now(),
			UpdatedBy: workerTag,
			UpdatedAt: gtime.Now(),
		})
		return nil
	})
	if err != nil {
		return 0, err
	}
	cache.id = workerId
	return cache.id, nil
}
