package threading

import (
	"github.com/hackfengJam/gokit/future"
	"github.com/hackfengJam/gokit/isafe"
)

// ThreadPoolExecutor Thread Pool Executor
type ThreadPoolExecutor struct {
	workerChan  chan future.PlaceholderType
	requestChan chan taskFunc
}

type taskFunc func()

// NewThreadPoolExecutor New ThreadPoolExecutor
func NewThreadPoolExecutor(maxWorkers int, maxWaiters int) *ThreadPoolExecutor {
	return &ThreadPoolExecutor{
		workerChan:  make(chan future.PlaceholderType, maxWorkers),
		requestChan: make(chan taskFunc, maxWaiters),
	}
}

// Submit submit
func (tpe *ThreadPoolExecutor) Submit(task taskFunc) {
	tpe.requestChan <- task
}

// workerReady Worker Ready
func (tpe *ThreadPoolExecutor) workerReady() {
	<-tpe.workerChan
}

// Run run
func (tpe *ThreadPoolExecutor) Run() {
	var taskQueue []taskFunc
	for {
		var req taskFunc
		select {
		case req = <-tpe.requestChan:
			taskQueue = append(taskQueue, req)
		case tpe.workerChan <- future.Placeholder:
			if len(taskQueue) == 0 {
				tpe.workerReady()
			} else {
				task := taskQueue[0]
				taskQueue = taskQueue[1:]
				go func() {
					defer isafe.Recover(func() {
						tpe.workerReady()
					})
					task()
				}()
			}
		}
	}
}

// Start start
func (tpe *ThreadPoolExecutor) Start() {
	go tpe.Run()
}
