package gjob

import (
	"sync"
)

var (
	instance *WorkerPool
	once     sync.Once
)

type Job func() // Job 定义一个无参数无返回值的函数类型，用于工作池中执行的任务。

// WorkerPool 结构体包含管理goroutine的工作通道和同步等待组。
type WorkerPool struct {
	JobChannel  chan Job       // JobChannel 是用于接收和执行作业的通道。
	stopChannel chan struct{}  // stopChannel 用于发送停止信号给工作goroutine。
	maxWorkers  int            // maxWorkers 表示工作池中的最大goroutine数量。
	wg          sync.WaitGroup // wg 用于同步等待所有goroutine完成。
}

func GetInstance(maxWorkers int) *WorkerPool {
	once.Do(func() {
		instance = &WorkerPool{
			JobChannel:  make(chan Job, maxWorkers),
			stopChannel: make(chan struct{}),
			maxWorkers:  maxWorkers,
		}
		instance.wg.Add(maxWorkers)
		for i := 0; i < maxWorkers; i++ {
			go instance.worker()
		}
	})
	return instance
}

// worker 是在goroutine中运行的函数，从JobChannel持续接收并执行作业。
func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for {
		select {
		case job := <-p.JobChannel: // 从JobChannel接收作业并执行
			job()
		case <-p.stopChannel: // 接收到停止信号，结束goroutine
			return
		}
	}
}

// Submit 将一个作业提交到工作池。
func (p *WorkerPool) Submit(job Job) {
	select {
	case p.JobChannel <- job:
		// 作业成功加入队列
	default:
		// 队列已满，启动一个临时worker来处理该作业
		go func() {
			job()
		}()
	}
}

// Close 关闭JobChannel并等待所有已经启动的作业完成。
func (p *WorkerPool) Close() {
	close(p.stopChannel) // 向所有worker发送停止信号
	p.wg.Wait()          // 等待所有worker完成
	close(p.JobChannel)  // 关闭作业通道
}

// IncreaseWorkers 增加指定数量的工作goroutine。
func (p *WorkerPool) IncreaseWorkers(num int) {
	p.wg.Add(num)
	for i := 0; i < num; i++ {
		go p.worker()
	}
}

// DecreaseWorkers 减少指定数量的工作goroutine。
func (p *WorkerPool) DecreaseWorkers(num int) {
	for i := 0; i < num; i++ {
		p.stopChannel <- struct{}{} // 发送停止信号给worker
	}
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
