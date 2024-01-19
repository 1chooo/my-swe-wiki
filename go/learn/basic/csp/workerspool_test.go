package csp_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
* Worker pool
 */
type TaskHandler func(interface{})

type Task struct {
	Param   interface{}
	Handler TaskHandler
}

type WorkerPoolImpl interface {
	AddWorker()    // 增加 worker
	SendTask(Task) // 發送任務
	Release()      // 釋放
}

type WorkerPool struct {
	wg   sync.WaitGroup
	inCh chan Task
}

func (d *WorkerPool) AddWorker() {
	d.wg.Add(1)
	go func() {
		for task := range d.inCh {
			task.Handler(task.Param)
		}
		d.wg.Done()
	}()
}

func (d *WorkerPool) Release() {
	close(d.inCh)
	d.wg.Wait()
}

func (d *WorkerPool) SendTask(t Task) {
	d.inCh <- t
}

func NewWorkerPool(buffer int) WorkerPoolImpl {
	return &WorkerPool{
		inCh: make(chan Task, buffer),
	}
}

func TestWorkersPool(t *testing.T) {
	bufferSize := 100
	var workerPool = NewWorkerPool(bufferSize)
	workers := 4
	for i := 0; i < workers; i++ {
		workerPool.AddWorker()
	}

	var sum int32
	testFunc := func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum, n)
	}
	var i, n int32
	n = 1000
	for ; i < n; i++ {
		task := Task{
			i,
			testFunc,
		}
		workerPool.SendTask(task)
	}
	workerPool.Release()
	fmt.Println(sum)
}

/*
協程池使用了反射來獲取執行的函數及參數。
如果批量執行的函數是已知的，可以優化成一種只執行指定函數的協程池，能夠提升性能。
*/
