package job

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gscheduler"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"strconv"
	"sync"
)

var (
	jobService service.IJobService
	once       sync.Once
)

type Service struct {
	task *gscheduler.TaskRunner
}

func New(task *gscheduler.TaskRunner) service.IJobService {
	once.Do(func() {
		jobService = &Service{
			task,
		}
		initTask(task)
	})
	return jobService
}

// 初始化定时任务
func initTask(task *gscheduler.TaskRunner) {
	sj := dal.SysJob
	//状态是重复执行的并且不是暂停的
	find, err := sj.WithContext(context.Background()).Where(sj.Status.Eq(0), sj.MisfirePolicy.Eq(1)).Find()
	if err != nil {
		gutils.ErrorPanic(err)
		return
	}
	for _, v := range find {
		err = task.AddTask(gscheduler.Task{
			ID:            int(v.ID),
			TaskFunName:   v.InvokeTarget,
			TaskFunParams: v.JobParams,
			ExecuteType:   v.MisfirePolicy,
			CronExpr:      v.CronExpression,
		})
		if err != nil {
			gutils.ErrorPanic(err)
			return
		}
	}
	task.Start()
	//订阅全局事件记录日志
	ghub.EventBus.Subscribe(constants.TaskLog, func(data interface{}) {
		TaskLog(data.(gscheduler.TaskResult))
	})
}

func TaskLog(res gscheduler.TaskResult) {
	var status int
	var errorStr = ""
	if res.Error != nil {
		ghub.Log.Info("定时任务执行失败")
		ghub.Log.Error(res.Error)
		status = 0
		errorStr = res.Error.Error()
	} else {
		ghub.Log.Info("定时任务执行成功")
		status = 1
	}
	ghub.Log.Info(res.ID)
	ghub.Log.Info(res.TaskFunName)
	ghub.Log.Info(res.TaskFunParams)
	ghub.Log.Info(res.CronExpr)
	ghub.Log.Info(res.ExecuteType)
	_ = dal.TaskLog.WithContext(context.Background()).Create(&model.TaskLog{
		JobId:    uint(res.ID),
		JobFunc:  res.TaskFunName,
		Status:   status,
		ErrorStr: errorStr,
	})
}

func (m *Service) Dao(ctx context.Context) dal.ISysJobDo {
	return dal.SysJob.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.JobAddReq) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		mModel := new(model.SysJob)
		if err := gutils.Copy(mModel, data); err != nil {
			return err
		}
		err := tx.SysJob.WithContext(ctx).Create(mModel)
		if err != nil {
			return err
		}
		return m.task.AddTask(gscheduler.Task{
			ID:            int(mModel.ID),
			TaskFunName:   mModel.InvokeTarget,
			TaskFunParams: mModel.JobParams,
			ExecuteType:   mModel.MisfirePolicy,
			CronExpr:      mModel.CronExpression,
		})
	})
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	dao := dal.SysJob
	info, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}

	//删除任务
	for _, id := range ids {
		m.task.PauseTask(int(id))
	}

	return nil
}

func (m *Service) DeleteJobLogAll(ctx context.Context, id uint) error {
	dao := dal.TaskLog
	info, err := dao.WithContext(ctx).Where(dao.JobId.Eq(id)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}
	return nil
}

func (m *Service) DeleteJobLog(ctx context.Context, ids []uint) error {
	dao := dal.TaskLog
	info, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}
	return nil
}

func (m *Service) Update(ctx context.Context, data *dto.JobUpdateReq) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.SysJob
		mModel, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).First()
		if err != nil {
			return fmt.Errorf(constants.NoDataFound)
		}
		if err = gutils.Copy(mModel, data); err != nil {
			return err
		}
		err = dao.WithContext(ctx).Save(mModel)
		if err != nil {
			return err
		}
		//暂停任务
		if mModel.Status == 1 {
			m.task.PauseTask(int(mModel.ID))
			return nil
		}
		//重置任务
		//删除任务
		m.task.RemoveTask(int(mModel.ID))
		//在启动任务
		return m.task.AddTask(gscheduler.Task{
			ID:            int(mModel.ID),
			TaskFunName:   mModel.InvokeTarget,
			TaskFunParams: mModel.JobParams,
			ExecuteType:   mModel.MisfirePolicy,
			CronExpr:      mModel.CronExpression,
		})
	})
}

func (m *Service) List(ctx context.Context, search *dto.JobSearchReq) ([]dto.Job, int64, error) {
	dao := dal.SysJob
	q := dao.WithContext(ctx)
	if search.JobName != "" {
		q = q.Where(dao.JobName.Eq(search.JobName))
	}
	if search.InvokeTarget != "" {
		q = q.Where(dao.InvokeTarget.Eq(search.InvokeTarget))
	}
	if search.JobGroup != "" {
		q = q.Where(dao.JobGroup.Eq(search.JobGroup))
	}
	if search.Status != "" {
		if status, err := strconv.Atoi(search.Status); err == nil {
			q = q.Where(dao.Status.Eq(status))
		}
	}

	list := make([]dto.Job, 0)
	count, err := q.Order(dao.ID.Asc()).ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (m *Service) JobLog(ctx context.Context, search *dto.JobLogReq) ([]dto.TaskLogRes, int64, error) {
	list := make([]dto.TaskLogRes, 0)
	dao := dal.TaskLog
	count, err := dao.WithContext(ctx).Where(dao.JobId.Eq(search.ID)).Order(dao.ID.Desc()).ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (m *Service) Once(ctx context.Context, id uint) error {
	dao := dal.SysJob
	job, err := dao.WithContext(ctx).Where(dao.ID.Eq(id)).First()
	if err != nil {
		return err
	}
	m.task.Once(gscheduler.Task{
		ID:            int(job.ID),
		TaskFunName:   job.InvokeTarget,
		TaskFunParams: job.JobParams,
		ExecuteType:   job.MisfirePolicy,
		CronExpr:      job.CronExpression,
	})
	return nil
}
