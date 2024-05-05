package gscheduler

import (
	"github.com/562589540/jono-gin/ghub/glibrary/geventbus"
	"github.com/562589540/jono-gin/ghub/glibrary/gjob"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

type ExecuteType int

const (
	Repeat ExecuteType = iota + 1
	Once
)

var (
	instance *TaskRunner
	once     sync.Once
)

// Task 代表数据库中存储的任务
type Task struct {
	ID            int         // 任务ID
	TaskFunName   string      // 方法名
	TaskFunParams string      // 任务需要的参数
	ExecuteType   ExecuteType // 执行类型：1表示重复执行，2表示执行一次
	CronExpr      string      // Cron 表达式
}

// TaskRunner 用于管理和执行所有的cron任务
type TaskRunner struct {
	Cron    *cron.Cron           // cron实例
	Entries map[int]cron.EntryID // 任务ID和cron.EntryID的映射
	lock    sync.RWMutex         // 添加读写锁
}

// cron.WithLogger(ghub.Log)

// GetInstance 返回TaskScheduler的单例实例。
func GetInstance() *TaskRunner {
	once.Do(func() {
		c := cron.New(cron.WithSeconds()) // 初始化cron实例，包括秒
		instance = &TaskRunner{
			Cron:    c,
			Entries: make(map[int]cron.EntryID),
		}
	})
	return instance
}

// Start 启动Cron服务
func (t *TaskRunner) Start() {
	t.Cron.Start()
}

// Stop 停止Cron服务
func (t *TaskRunner) Stop() {
	t.Cron.Stop()
}

// AddTask 添加新任务
func (t *TaskRunner) AddTask(task Task) error {
	t.lock.Lock() // 写锁定
	defer t.lock.Unlock()
	var id cron.EntryID
	var err error
	if task.ExecuteType == Repeat { // 重复执行
		id, err = t.Cron.AddFunc(task.CronExpr, func() {
			executeTask(task) // 执行任务的逻辑
		})
	} else if task.ExecuteType == Once { // 执行一次
		id, err = t.Cron.AddFunc(task.CronExpr, func() {
			executeTask(task)     // 执行任务的逻辑
			t.RemoveTask(task.ID) // 执行后立即删除任务
		})
	}
	if err != nil {
		log.Println(err.Error())
		return err
	}
	t.Entries[task.ID] = id
	return nil
}

// RemoveTask 删除任务
func (t *TaskRunner) RemoveTask(taskID int) {
	t.lock.Lock() // 写锁定
	defer t.lock.Unlock()
	if entryID, ok := t.Entries[taskID]; ok {
		t.Cron.Remove(entryID)
		delete(t.Entries, taskID)
	}
}

// PauseTask 暂停任务
func (t *TaskRunner) PauseTask(taskID int) {
	// 暂停实际上是移除现有任务
	t.RemoveTask(taskID)
}

// ResumeTask 恢复任务
func (t *TaskRunner) ResumeTask(task Task) error {
	// 恢复任务实际上是重新添加任务
	err := t.AddTask(task)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Once 执行一次
func (t *TaskRunner) Once(task Task) {
	executeTask(task)
}

type TaskResult struct {
	Task
	Error error
}

// executeTask 执行具体的任务
func executeTask(task Task) {
	// 这里应该根据 task.TaskFunName 和 task.TaskFunParams 调用相应的方法
	gjob.GetInstance(100).Submit(func() {
		err := callFunc(task.TaskFunName, task.TaskFunParams)
		geventbus.GetInstance().Publish(constants.TaskLog, TaskResult{
			Task:  task,
			Error: err,
		})
	})
}

// IsValidCronExpression 检查给定的 Cron 表达式是否有效
func IsValidCronExpression(expr string) bool {
	// 使用带有所有选项的解析器，包括秒
	parser := cron.NewParser(
		cron.Second | // 支持秒
			cron.Minute | // 支持分钟
			cron.Hour | // 支持小时
			cron.Dom | // 支持月中的某天
			cron.Month | // 支持月份
			cron.DowOptional | // 支持可选的星期几
			cron.Descriptor, // 支持特定描述符
	)
	if _, err := parser.Parse(expr); err != nil {
		return false
	}
	return true
}

//
//type Config struct {
//	WithSeconds bool
//}
//
//func NewTaskRunner(config *Config) *TaskRunner {
//	var options []cron.Option
//	if config.WithSeconds {
//		options = append(options, cron.WithSeconds())
//	}
//	c := cron.New(options...)
//	return &TaskRunner{
//		Cron:    c,
//		Entries: make(map[int]cron.EntryID),
//	}
//}
