package guid

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"

	_ "github.com/chaos-plus/chaos-plus/internal/packed"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/configs"
)

type sGuid struct{}

func init() {
	service.RegisterGuid(&sGuid{})
}

func New() service.IGuid {
	return &sGuid{}
}

func (s *sGuid) InitOrPanic(ctx context.Context) {
	basetime := configs.GetConfigOrDefault(ctx, "guid.basetime", "2025-01-01 00:00:00").String()
	workerid := service.WorkerId().GetWorkerId(ctx)
	options := idgen.NewIdGeneratorOptions(uint16(workerid))
	options.WorkerIdBitLength = 10                     // 默认值6,限定 WorkerId 最大值为2^6-1,即默认最多支持64个节点。
	options.SeqBitLength = 10                          // 默认值6,限制每毫秒生成的ID个数。若生成速度超过5万个/秒,建议加大 SeqBitLength 到 10。
	options.BaseTime = gtime.New(basetime).UnixMilli() // 如果要兼容老系统的雪花算法,此处应设置为老系统的BaseTime。
	// ...... 其它参数参考 IdGeneratorOptions 定义。
	// 保存参数（务必调用,否则参数设置不生效）:
	idgen.SetIdGenerator(options)
}

func (s *sGuid) NextId() uint64 {
	return uint64(idgen.NextId())
}
