package gscheduler

import (
	"fmt"
	"sync"
	"time"
)

type TaskState string

const (
	Waiting   TaskState = "waiting"
	Running   TaskState = "running"
	Completed TaskState = "completed"
	Cancelled TaskState = "cancelled"
)

// TaskFunc 定义任务执行的函数类型。
type TaskFunc func()

// ScheduledTask 表示一个计划中的任务。
type ScheduledTask struct {
	ID       string
	Interval time.Duration
	Task     TaskFunc
	Repeat   bool
	ticker   *time.Ticker
	quit     chan struct{}
	State    TaskState
}

// TaskScheduler 管理所有定时任务的调度器。
type TaskScheduler struct {
	tasks map[string]*ScheduledTask
	mu    sync.Mutex
}

var (
	instance *TaskScheduler
	once     sync.Once
)

// GetInstance 返回TaskScheduler的单例实例。
func GetInstance() *TaskScheduler {
	once.Do(func() {
		instance = &TaskScheduler{
			tasks: make(map[string]*ScheduledTask),
		}
	})
	return instance
}

// ScheduleTask 添加并启动一个定时任务。任务可以是重复的或一次性的。
func (s *TaskScheduler) ScheduleTask(id string, interval time.Duration, task TaskFunc, repeat bool) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保任务ID是唯一的
	if _, exists := s.tasks[id]; exists {
		return id // 如果ID已存在，直接返回
	}

	newTask := &ScheduledTask{
		ID:       id,
		Interval: interval,
		Task:     task,
		Repeat:   repeat,
		quit:     make(chan struct{}),
		State:    Waiting,
	}

	if repeat {
		newTask.ticker = time.NewTicker(newTask.Interval)
		go func() {
			newTask.State = Running
			defer func() {
				newTask.State = Completed
				if r := recover(); r != nil {
					fmt.Println("Recovered in task:", r)
					newTask.State = Cancelled
				}
			}()
			for {
				select {
				case <-newTask.ticker.C:
					newTask.Task()
				case <-newTask.quit:
					newTask.ticker.Stop()
					newTask.State = Cancelled
					return
				}
			}
		}()
	} else {
		// 一次性定时任务完成后自动移除
		time.AfterFunc(newTask.Interval, func() {
			newTask.State = Running
			newTask.Task()
			newTask.State = Completed
			s.RemoveTask(newTask.ID)
		})
	}

	s.tasks[id] = newTask
	return id
}

// StopTask 停止并移除指定ID的任务。
func (s *TaskScheduler) StopTask(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, exists := s.tasks[id]; exists {
		close(task.quit)
		delete(s.tasks, id)
	}
}

// RemoveTask 从调度器中移除任务。
func (s *TaskScheduler) RemoveTask(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; exists {
		delete(s.tasks, id)
	}
}

// StopAll 停止并移除所有任务。
func (s *TaskScheduler) StopAll() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, task := range s.tasks {
		close(task.quit)
		delete(s.tasks, id)
	}
}

// GetTaskState 返回指定任务的状态。
func (s *TaskScheduler) GetTaskState(id string) (TaskState, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if task, exists := s.tasks[id]; exists {
		return task.State, true
	}
	return "", false // 返回空字符串表示未找到任务
}

//s := gscheduler.GetInstance()
//
//taskID := s.ScheduleTask("task1", 2*time.Second, func() {
//    fmt.Println("Task 1 executed")
//}, true)
//
//// 获取并打印任务状态
//state, _ := s.GetTaskState(taskID)
//fmt.Println("Task State:", state)
//
//// 运行一段时间后停止任务，并检查状态
//time.Sleep(5 * time.Second)
//s.StopTask(taskID)
//state, _ = s.GetTaskState(taskID)
//fmt.Println("Task State after stopping:", state)
