package jobs

import (
	"fmt"
	"time"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/sev-2/raiden"
)

type HelloWorldJob struct {
	raiden.JobBase
}

func (j *HelloWorldJob) Name() string {
	return "hello_world_job"
}

func (j *HelloWorldJob) Before(ctx raiden.JobContext, jobID uuid.UUID, jobName string) {
	raiden.Info("before execute", "name", jobName)
}

func (j *HelloWorldJob) After(ctx raiden.JobContext, jobID uuid.UUID, jobName string) {
	raiden.Info("after execute", "name", jobName)
}

func (j *HelloWorldJob) AfterErr(ctx raiden.JobContext, jobID uuid.UUID, jobName string, err error) {
	raiden.Error("after execute with error", "message", err.Error())
}

func (j *HelloWorldJob) Duration() gocron.JobDefinition {
	return gocron.DurationJob(5 * time.Minute)
}

func (j *HelloWorldJob) Task(ctx raiden.JobContext) error {
	fmt.Printf("Hello at %s \n", time.Now().String())

	return nil
}
