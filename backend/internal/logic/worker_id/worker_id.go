package workerid

import (
	"context"
	"time"

	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/crypto"
	netutils "github.com/chaos-plus/chaos-plus/utility/net"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/shirou/gopsutil/v3/host"
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
	// service.RegisterWorkerId(&sWorkerId{})
}

// func New() service.IWorkerId {
// 	return &sWorkerId{}
// }

func RefreshWorkerId() {
}

func (s *sWorkerId) getExists(ctx context.Context, hostTag string) (res *entity.WorkerId, err error) {
	m := dao.WorkerId.Ctx(ctx)
	err = m.Where("host_tag = ?", hostTag).Scan(&res)
	return
}

func (s *sWorkerId) getExpired(ctx context.Context, expiredTime int64) (res *entity.WorkerId, err error) {
	m := dao.WorkerId.Ctx(ctx)
	err = m.Where("time_created < ?", time.Now().Unix()-expiredTime).Limit(1).Scan(&res)
	return
}

// 获取 workerId
func (s *sWorkerId) getWorkerId(ctx context.Context) (uint, error) {
	if cache.id != 0 {
		return cache.id, nil
	}

	hostInfo, _ := host.Info()

	workerInfo := gjson.New(make(map[string]interface{}))
	workerInfo.Set("host", hostInfo)
	workerInfo.Set("nets", netutils.GetLanAll())

	workerId := &do.WorkerId{
		Id:       1,
		HostInfo: workerInfo.String(),
		HostTag:  crypto.Md5(workerInfo.String()),
	}

	dao.WorkerId.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		tx.Model(&entity.WorkerId{}).Insert(workerId)
		return nil
	})

	// get distributed locker

	// delete expired workerId

	// get all workerId

	// generate new workerId

	// add workerId

	// release distributed locker

	cache.id = workerId.Id.(uint)
	return cache.id, nil
}
