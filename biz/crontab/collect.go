package crontab

import "github.com/CloudGoSight/cloudgosight_api/util"

func garbageCollect() {
	// 清理打包下载产生的临时文件
	//collectArchiveFile()

	// 清理过期的内置内存缓存
	//if store, ok := cache.Store.(*cache.MemoStore); ok {
	//	collectCache(store)
	//}

	util.Log().Info("Crontab job \"cron_garbage_collect\" complete.")
}

func uploadSessionCollect() {
	util.Log().Info("Crontab job \"upload_session_collect\" complete.")
}
