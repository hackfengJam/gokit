package filter

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/hackfengJam/filter/core"
	"github.com/hackfengJam/gokit/cast"
)

var (
	alwaysTrueCondition core.Condition
)

func init() {
	var err error
	alwaysTrueCondition, err = NewConditionByJSONString(`["succ", "=", true]`, core.LOGIC_ALL)
	if err != nil {
		// 解析条件异常。报错
		glog.Fatal(err)
		return
	}
}

// GetAlwaysTrueCondition Get AlwaysTrueCondition
func GetAlwaysTrueCondition() core.Condition {
	return alwaysTrueCondition
}

// NewConditionByJSONString New Condition By JSONString
func NewConditionByJSONString(jsonStr string, groupLogic core.GROUP_LOGIC) (core.Condition, error) {
	var filterData []interface{}
	if err := json.Unmarshal(cast.StringToBytes(jsonStr), &filterData); err != nil {
		return nil, err
	}
	return core.NewCondition(filterData, groupLogic)
}

// NewCondition New Condition
func NewCondition(item []interface{}, groupLogic core.GROUP_LOGIC) (core.Condition, error) {
	return core.NewCondition(item, groupLogic)
}
