package isprint

import (
	"sync"
	"testing"
)

func Test_FormatSyncMap(t *testing.T) {
	type Std struct {
		Name int64
		Age  int64
	}

	var data sync.Map
	data.Store("SiteData", map[string]interface{}{
		"Name": "StackOverflow",
		"Url":  "https://so.com",
		"Std":  &Std{Name: 1, Age: 2},
	})
	data.Store("Else", "something else")
	data.Store(111, &Std{Name: 3, Age: 2})

	tests := []struct {
		obj *sync.Map
		ret string
	}{
		{
			obj: &data,
			ret: "",
		},
		{
			obj: nil,
			ret: "",
		},
	}

	for _, test := range tests {
		t.Log(FormatSyncMap(test.obj))
	}
}

func Benchmark_FormatSyncMap(b *testing.B) {
	type Std struct {
		Name int64
		Age  int64
	}

	var data sync.Map
	data.Store("SiteData", map[string]interface{}{
		"Name": "StackOverflow",
		"Url":  "https://so.com",
		"Std":  &Std{Name: 1, Age: 2},
	})
	data.Store("Else", "something else")
	data.Store(111, &Std{Name: 3, Age: 2})

	for i := 1; i < b.N; i++ {
		FormatSyncMap(&data)
	}
}
