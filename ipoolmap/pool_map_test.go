package ipoolmap

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoolMapOrder(t *testing.T) {
	funcs := make([]func() (interface{}, error), 0, 2)
	for i := 0; i < 6; i++ {
		var j int
		j = i
		f := func() (interface{}, error) {
			return j, nil
		}
		funcs = append(funcs, f)
	}

	result, err := PoolMap(context.Background(), funcs)

	assert.Equal(t, err, nil)
	for i := 0; i < 6; i++ {
		assert.EqualValues(t, fmt.Sprint(result[i].Data), fmt.Sprint(i))
	}
}

func TestPoolMapPanic(t *testing.T) {
	funcs := make([]func() (interface{}, error), 0, 2)
	for i := 0; i < 6; i++ {
		f := func() (interface{}, error) {
			result := []int64{1, 2}
			// make a slice bound panic.
			return result[4], nil
		}
		funcs = append(funcs, f)
	}

	_, err := PoolMap(context.Background(), funcs)
	assert.EqualValues(t, err, ErrPoolMapExecute)
}

func TestPoolMapErr(t *testing.T) {
	func1 := func() (interface{}, error) {
		return true, nil
	}
	func2 := func() (interface{}, error) {
		return nil, errors.New("test")
	}
	result, _ := PoolMap(context.Background(), []func() (interface{}, error){func1, func2})
	assert.EqualValues(t, result[0].Data.(bool), true)
	assert.EqualValues(t, result[1].Err, errors.New("test"))
}
