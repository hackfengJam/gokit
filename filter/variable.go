package filter

import (
	"github.com/hackfengJam/filter/core"
)

func init() {
}

// RegisterFilterDataKeyValue 带有默认值的 filter.RegisterVariableFunc
// RegisterFilterDataKeyValue 把变量（keys[0]）注入到 filter
// keys 表示[自定义变量，属性key]，若 2 者一致可以只传入 1 个
func RegisterFilterDataKeyValue(defaultValue interface{}, keys ...string) {
	var key, mapkey string
	if len(keys) == 1 {
		key = keys[0]
		mapkey = key
	} else {
		key = keys[0]
		mapkey = keys[1]
	}
	RegisterVariableFunc(key, func(ctx *core.Context) interface{} {
		if d, ok := ctx.Data().(map[string]interface{}); ok {
			if value, iOk := d[mapkey]; iOk {
				return value
			}
		}
		return defaultValue
	}, false)
}

// RegisterVariableFunc Register VariableFunc
// 在启动的时候注册，业务方可以注入自己的 variable 解析函数
// 在 GetVariableValue 的时候可以获取 key，方便 debug
// 除非自己指定不使用缓存，其余默认使用缓存
func RegisterVariableFunc(name string, f core.ValueFunc, cacheable bool) {
	core.GetVariableFactory().Register(core.SingletonVariableCreator(core.NewSimpleVariable(name, cacheable, f)), name)
}

// CreateVariable create variable
func CreateVariable(name string) core.Variable {
	return core.GetVariableFactory().Create(name)
}
