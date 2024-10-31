package crontab

import (
	"github.com/CloudGoSight/cloudgosight_api/util"
	"github.com/robfig/cron/v3"
)

// Init 初始化定时任务
func Init() {
	util.Log().Info("Initialize crontab jobs...")
	// 读取cron日程设置
	options := make(map[string]string, 2)
	options["cron_garbage_collect"] = "@hourly"
	options["cron_recycle_upload_session"] = "@every 1h30m"
	Cron := cron.New()
	for k, v := range options {
		var handler func()
		switch k {
		case "cron_garbage_collect":
			handler = garbageCollect
		case "cron_recycle_upload_session":
			handler = uploadSessionCollect
		default:
			util.Log().Info("Unknown crontab job type %q, skipping...", k)
			continue
		}

		if _, err := Cron.AddFunc(v, handler); err != nil {
			util.Log().Info("Failed to start crontab job %q: %s", k, err)
		}

	}
	Cron.Start()
}
