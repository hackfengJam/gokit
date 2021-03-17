package threading

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThreadPoolExecutor(t *testing.T) {
	times := 10000
	executor := NewThreadPoolExecutor(runtime.NumCPU(), 10000000)
	// pool := NewThreadPoolExecutor(3, 10000000)
	start := time.Now()
	var counter int32
	var waitGroup sync.WaitGroup
	for i := 0; i < times; i++ {
		waitGroup.Add(1)
		executor.Submit(func() {
			time.Sleep(time.Millisecond)
			atomic.AddInt32(&counter, 1)
			waitGroup.Done()
		})
		// fmt.Println(fmt.Sprintf("th: %d, has been submitted", i))
	}

	fmt.Println("submit finished.")
	waitGroup.Wait()
	fmt.Println(fmt.Sprintf("elapsed: %v", time.Now().Sub(start)))

	assert.Equal(t, times, int(counter))
}
