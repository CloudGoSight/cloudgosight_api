package aria2

import (
	"github.com/CloudGoSight/cloudgosight_api/biz/balance"
	"github.com/CloudGoSight/cloudgosight_api/biz/cluster"
	"github.com/CloudGoSight/cloudgosight_api/biz/dal/mq"
	"sync"
)

// LB 获取 Aria2 节点的负载均衡器
var LB balance.Balancer

// Lock Instance的读写锁
// RWMutex 也称为读写互斥锁，读写互斥锁就是读取/写入互相排斥的锁。
// 它可以由任意数量的读取操作的 goroutine 或单个写入操作的 goroutine 持有
var Lock sync.RWMutex

func Init(isReload bool, pool cluster.Pool, mqClient mq.MQ) {
	Lock.Lock()
	LB = balance.NewBalancer("round_robin")
	Lock.Unlock()

	//if !isReload {
	//	// 从数据库中读取未完成任务，创建监控
	//	unfinished := model.GetDownloadsByStatus(common.Ready, common.Paused, common.Downloading, common.Seeding)
	//
	//	for i := 0; i < len(unfinished); i++ {
	//		// 创建任务监控
	//		monitor.NewMonitor(&unfinished[i], pool, mqClient)
	//	}
	//}
}
