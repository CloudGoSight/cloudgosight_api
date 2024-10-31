package bootstrap

import (
	"github.com/CloudGoSight/cloudgosight_api/biz/aria2"
	"github.com/CloudGoSight/cloudgosight_api/biz/crontab"
	"github.com/CloudGoSight/cloudgosight_api/biz/dal/mq"
	"github.com/CloudGoSight/cloudgosight_api/biz/dal/redis"
	"github.com/CloudGoSight/cloudgosight_api/biz/task"
	"github.com/CloudGoSight/cloudgosight_api/conf"
	"io/fs"
)

func Init(path string, statics fs.FS) {
	InitApplication()
	conf.Init(path)

	dependencies := []struct {
		mode    string
		factory func()
	}{
		{
			"both",
			func() {
				redis.Init()
			},
		},
		{
			"master",
			func() {
				// mysql
			},
		},
		{
			"both",
			func() {
				task.Init()
			},
		},
		{
			"master",
			func() {
				// cluster
			},
		},
		{
			"master",
			func() {
				aria2.Init(false, nil, mq.GlobalMQ)
			},
		},
		{
			"master",
			func() {
				// email
			},
		},
		{
			"master",
			func() {
				crontab.Init()
			},
		},
		{
			"master",
			func() {
				// statics
			},
		},
		{
			"slave",
			func() {
				// controller
			},
		},
		{
			"both",
			func() {
				// auth
			},
		},
		{
			"master",
			func() {
				// WOPI
			},
		},
	}

	for _, dependency := range dependencies {
		if dependency.mode == conf.SystemConfig.Mode || dependency.mode == "both" {
			dependency.factory()
		}
	}
}
