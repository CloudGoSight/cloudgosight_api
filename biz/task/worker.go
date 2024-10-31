package task

import (
	"fmt"
	"github.com/CloudGoSight/cloudgosight_api/util"
)

type Worker interface {
	Do(job Job)
}

type GeneralWorker struct{}

func (g GeneralWorker) Do(job Job) {
	util.Log().Info("start executing job")
	job.SetStatus(Processing)

	defer func() {
		if err := recover(); err != nil {
			util.Log().Info("recovered from panic %s", err)
			job.SetError(&JobError{Msg: "recovered from panic", Err: fmt.Sprintf("%s", err)})
			job.SetStatus(Error)
		}
	}()

	job.Do()

	if err := job.GetError(); err != nil {
		util.Log().Info("job get err %s", err)
		job.SetStatus(Error)
		return
	}

	util.Log().Info("job finished successfully")
	job.SetStatus(Complete)
}
