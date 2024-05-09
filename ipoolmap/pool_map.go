package ipoolmap

import (
	"context"
	"errors"
	"sort"
	"sync"

	"github.com/hackfengJam/gokit/isafe"
)

// 并发执行一组函数，并将结果返回。
// ⚠️：没有处理函数执行的超时，调用者需要自己控制超时。
// TODO 支持控制并发数量。
// map 命名参考 函数式编程的 map reduce 的命名。
// PoolMap 使用场景：适用于小批量并发请求。

// PoolResult Pool Result
type PoolResult struct {
	idx  int
	Data interface{}
	Err  error
}

// var
var (
	ErrPoolMapExecute = errors.New("pool map execute Err")
)

// PoolMap Pool Map
/*
example:
	type TestResult struct {
        UID int64
	}
	args := []int64{1,2,3}

	// 定义执行函数
	funcs := make([]func() (interface{}, error), 0)
	for _, v := range args {
		uid := v
		funcs = append(funcs, func() (interface{}, error) {
			// do something
			item := &TestResult{}
			item.UID = uid
			return item, nil
		})
	}

	// 并发执行&等待所有结果
	results, err := ipoolmap.PoolMap(ctx, funcs)
	if err != nil {
		// log
		return
	}

	// 获取结果
	for _, result := range results {
		if result.Err != nil {
			continue
		}
		item, ok := result.Data.(*TestResult)
		if !ok {
			continue
		}
		resultList = append(resultList, item)
	}
*/
func PoolMap(ctx context.Context, funcs []func() (interface{}, error)) ([]PoolResult, error) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	resultChan := make(chan PoolResult, len(funcs))
	for i, fn := range funcs {
		go func(idx int, f func() (interface{}, error)) {
			defer isafe.Recover()
			defer wg.Done()
			result, err := f()
			resultChan <- PoolResult{
				idx:  idx, // 用 idx 维护是传入函数的结果列表。
				Data: result,
				Err:  err,
			}
		}(i, fn)
	}
	retResults := make([]PoolResult, 0, len(funcs))
	wg.Wait()
	close(resultChan)
	for result := range resultChan {
		retResults = append(retResults, result)
	}
	if len(retResults) < len(funcs) {
		// 至少有一个 goroutine 发生了panic。 在 goroutine 里出现 panic，没有结果。
		return nil, ErrPoolMapExecute
	}
	// 将执行结果排序为传入的顺序。
	sort.Slice(retResults, func(i, j int) bool {
		return retResults[i].idx < retResults[j].idx
	})
	return retResults, nil
}
