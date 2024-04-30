package gjob

import (
	"sync"
)

var (
	instance *WorkerPool
	once     sync.Once
)

// Job 定义了一个无参数无返回值的函数类型，用于工作池中的任务。
type Job func()

// WorkerPool 结构包含了管理goroutine的工作通道和同步等待组。
type WorkerPool struct {
	JobChannel chan Job       // JobChannel 是用于接收和执行作业的通道。
	wg         sync.WaitGroup // wg 用于同步等待所有goroutine完成。
}

// GetInstance 返回一个WorkerPool的单例实例。
// workerCount 是启动时的工作goroutine数量。
// 如果实例尚未创建，它会初始化一个具有指定数量worker的WorkerPool。
func GetInstance(workerCount int) *WorkerPool {
	once.Do(func() {
		instance = &WorkerPool{
			JobChannel: make(chan Job),
		}
		instance.wg.Add(workerCount)
		for i := 0; i < workerCount; i++ {
			go instance.worker()
		}
	})
	return instance
}

// worker 是在goroutine中运行的函数，从JobChannel持续接收并执行作业。
func (p *WorkerPool) worker() {
	for job := range p.JobChannel {
		job()
	}
	p.wg.Done()
}

// Submit 将一个作业提交到工作池。
func (p *WorkerPool) Submit(job Job) {
	p.JobChannel <- job
}

// Close 关闭JobChannel并等待所有已经启动的作业完成。
func (p *WorkerPool) Close() {
	close(p.JobChannel)
	p.wg.Wait()
}

//pool := gjob.GetWorkerPool(5) // 初始化或获取已有的WorkerPool实例
//	pool.Submit(func() {
//		fmt.Println("Running gjob 1")
//		time.Sleep(1 * time.Second)
//	})
//	pool.Submit(func() {
//		fmt.Println("Running gjob 2")
//		time.Sleep(1 * time.Second)
//	})
//	// 当所有作业完成时关闭工作池
//	defer pool.Close()
