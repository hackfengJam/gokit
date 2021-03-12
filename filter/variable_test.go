package filter

import (
	"context"
	"testing"

	"github.com/hackfengJam/filter/core"
	"github.com/stretchr/testify/assert"
)

func TestCreateVariable(t *testing.T) {
	ast := assert.New(t)

	ctx := core.WithContext(context.Background())
	ctx = core.WithData(ctx, map[string]interface{}{
		"ctx": map[string]interface{}{
			"baz": "baz-in-data",
		},
	})
	ctx.Set("foo", map[string]interface{}{
		"bar": "bar-ctx",
	})
	ctx.Set("baz", "baz-ctx")

	tests := []struct {
		input    string
		expected interface{}
	}{
		{"ctx.foo.bar", "bar-ctx"},
		{"ctx.baz", "baz-in-data"},
		{"ctx.other", nil},
	}

	for i, c := range tests {
		v := CreateVariable(c.input)
		ast.NotNil(v, c.input)
		ast.Equal(c.expected, core.GetVariableValue(ctx, v), "case %d : %s", i, c.input)
	}
}

func TestRegisterVariableFunc(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		input    string
		expected interface{}
	}{
		{"a", "value_a"},
		{"a.b", "value_a_b"},
		{"a.c", "value_a_c"},
		{"succ", true},
	}
	for _, c := range tests {
		func(name string, value interface{}) {
			RegisterVariableFunc(name, func(ctx *core.Context) interface{} {
				return value
			}, false)
		}(c.input, c.expected)
	}

	ctx := core.WithContext(context.Background())
	for i, c := range tests {
		v := CreateVariable(c.input)
		ast.NotNil(v, c.input)
		ast.Equal(c.expected, core.GetVariableValue(ctx, v), "case %d : %s", i, c.input)
	}
}
