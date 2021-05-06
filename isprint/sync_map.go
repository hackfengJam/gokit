package isprint

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/hackfengJam/gokit/cast"
)

// FormatSyncMap Format SyncMap
func FormatSyncMap(data *sync.Map) string {
	if data == nil {
		return ""
	}

	m := map[string]interface{}{}
	data.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = value
		return true
	})

	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return ""
	}
	return string(cast.BytesToString(b))
}
