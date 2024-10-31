package task

import (
	"github.com/CloudGoSight/cloudgosight_api/conf"
	"github.com/CloudGoSight/cloudgosight_api/util"
)

type Pool interface {
	Add(num int)
	Submit(job Job)
}

type AsyncPool struct {
	// 容量
	idleWorker chan int
}

var TaskPool Pool

// 增加可用worker数量
func (pool *AsyncPool) Add(num int) {
	for i := 0; i < num; i++ {
		pool.idleWorker <- 1
	}
}

// 阻塞直到获取新的worker
func (pool *AsyncPool) obtainWorker() Worker {
	select {
	case <-pool.idleWorker:
		return &GeneralWorker{}
	}
}

// 添加空闲worker
func (pool *AsyncPool) freeWorker() {
	pool.Add(1)
}

func (pool *AsyncPool) Submit(job Job) {
	go func() {
		util.Log().Info("Waiting for worker")
		worker := pool.obtainWorker()
		util.Log().Info("Worker obtained")
		worker.Do(job)
		util.Log().Info("Worker released")
		pool.freeWorker()
	}()
}

func Init() {
	maxWorker := 10 //todo:从redis+mysql取
	TaskPool = &AsyncPool{make(chan int, maxWorker)}
	TaskPool.Add(maxWorker)
	util.Log().Info("Init pool")
	if conf.SystemConfig.Mode == "master" {
		// todo:resume
	}
}
