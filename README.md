# gokit
A common library for golang.

## Installation

```
go get github.com/hackfengJam/gokit
```

## Quick Start

### ipoolmap

```go
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
        // do something
        continue
    }
    item, ok := result.Data.(*TestResult)
    if !ok {
        // do something
        continue
    }
    resultList = append(resultList, item)
}
```

8. Documents (adding)

- TODO

## Give a Star! 

If you like or are using this project to learn or start your solution, please give it a star. Thanks!