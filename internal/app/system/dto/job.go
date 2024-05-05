package dto

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gscheduler"
	"time"
)

type JobLogReq struct {
	ID               uint `json:"id" form:"id" binding:"required"`
	gdto.PaginateReq      //分页
}

type TaskLogRes struct {
	ID        uint      `json:"id"`
	JobId     uint      `json:"jobId"`
	JobFunc   string    `json:"jobFunc"`
	ErrorStr  string    `json:"errorStr"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type JobSearchReq struct {
	gdto.PaginateReq        //分页
	JobName          string `json:"jobName" form:"jobName"`           // 任务名称,
	InvokeTarget     string `json:"invokeTarget" form:"invokeTarget"` // 任务方法,
	JobGroup         string `json:"jobGroup" form:"jobGroup"`         // 任务组名,
	Status           string `json:"status" form:"status"`             // 状态,

}

type JobUpdateReq struct {
	ID uint `json:"id" binding:"required"`
	JobAddReq
}

type JobAddReq struct {
	CronExpression string                 `json:"cronExpression" binding:"required" m:"cron执行表达式不能为空"` // cron执行表达式,
	JobName        string                 `json:"jobName" binding:"required" m:"任务名称不能为空"`             // 任务名称,
	InvokeTarget   string                 `json:"invokeTarget" binding:"required" m:"任务方法不能为空"`        // 任务方法,
	JobParams      string                 `json:"jobParams"`                                           // 任务参数,
	JobGroup       string                 `json:"jobGroup"`                                            // 任务组名,
	MisfirePolicy  gscheduler.ExecuteType `json:"misfirePolicy"`                                       // 计划执行策略,
	Status         int                    `json:"status"`                                              // 状态,
	Remark         string                 `json:"remark"`                                              // 备注
}

func (dto JobAddReq) Validate() error {
	if !gscheduler.IsValidCronExpression(dto.CronExpression) {
		return fmt.Errorf("cron 表达式非法")
	}
	return nil
}

type Job struct {
	ID             uint   `json:"id"`             // ,
	CronExpression string `json:"cronExpression"` // cron执行表达式,
	JobName        string `json:"jobName"`        // 任务名称,
	InvokeTarget   string `json:"invokeTarget"`   // 任务方法,
	JobParams      string `json:"jobParams"`      // 任务参数,
	JobGroup       string `json:"jobGroup"`       // 任务组名,
	MisfirePolicy  int    `json:"misfirePolicy"`  // 计划执行策略,
	Status         int    `json:"status"`         // 状态,
	Remark         string `json:"remark"`         // 备注
}
