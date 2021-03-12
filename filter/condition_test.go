package filter

import (
	"encoding/json"
	"testing"

	"github.com/hackfengJam/filter/core"
	"github.com/stretchr/testify/assert"
)

func TestNewCondition(t *testing.T) {
	ctx := core.NewContext()
	ctx.Set("a", map[string]interface{}{
		"b": 1,
		"c": []interface{}{1, 2},
		"d": 5,
	})

	tests := []struct {
		item     []interface{}
		logic    core.GROUP_LOGIC
		expected bool
		hasError bool
	}{
		{
			item:     []interface{}{"ctx.a.b", "=", 1},
			logic:    core.LOGIC_ALL,
			expected: true,
		},
		{
			item:     []interface{}{"ctx.a.b", "=", 2},
			logic:    core.LOGIC_ALL,
			expected: false,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 1},
				[]interface{}{"ctx.a.c.0", "=", 1},
				[]interface{}{"ctx.a.d", "=", 5},
			},
			logic:    core.LOGIC_ALL,
			expected: true,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 2},
				[]interface{}{"ctx.a.c.0", "=", 1},
				[]interface{}{"ctx.a.d", "=", 5},
			},
			logic:    core.LOGIC_ALL,
			expected: false,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 2},
				[]interface{}{"ctx.a.c.0", "=", 1},
				[]interface{}{"ctx.a.d", "=", 5},
			},
			logic:    core.LOGIC_ANY,
			expected: true,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 2},
				[]interface{}{"ctx.a.c.0", "=", 22},
				[]interface{}{"ctx.a.d", "=", 55},
			},
			logic:    core.LOGIC_ANY,
			expected: false,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 1},
				[]interface{}{"ctx.a.c.0", "=", 1},
			},
			logic:    core.LOGIC_NONE,
			expected: false,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 2},
			},
			logic:    core.LOGIC_NONE,
			expected: true,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 2},
				[]interface{}{"ctx.a.c.0", "=", 1},
			},
			logic:    core.LOGIC_NONE,
			expected: false,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 2},
				[]interface{}{"ctx.a.c.0", "=", 1},
			},
			logic:    core.LOGIC_ANY_NOT,
			expected: true,
		},
		{
			item: []interface{}{
				[]interface{}{"ctx.a.b", "=", 1},
				[]interface{}{"ctx.a.c.0", "=", 1},
			},
			logic:    core.LOGIC_ANY_NOT,
			expected: false,
		},
		{
			item: []interface{}{
				"any?", "=>", []interface{}{
					[]interface{}{"ctx.a.b", "=", 2},
					[]interface{}{
						"none?", "=>", []interface{}{
							[]interface{}{"ctx.a.c.0", "=", 2},
						},
					},
				},
			},
			logic:    core.LOGIC_ALL,
			expected: true,
		},
	}

	ast := assert.New(t)

	for i, c := range tests {
		cond, err := NewCondition(c.item, c.logic)
		if c.hasError {
			ast.Error(err)
			continue
		}
		ast.NoError(err, err)
		ast.Equal(c.expected, cond.Success(ctx), "case %d: %s", i, cond.String())
	}
}

func TestBiz(t *testing.T) {
	ast := assert.New(t)

	ctx := core.NewContext()

	jsonStr := `[["has_magic_no_list", "has", "214"],["is_finished_guide_step", "=", true]]`
	// jsonStr = `[["succ", "=", 1]]`
	// jsonStr = `["succ", "=", 1]`
	var filterData []interface{}
	if err := json.Unmarshal([]byte(jsonStr), &filterData); err != nil {
		t.Errorf("test case unmarshal json fail:%s", err)
	}

	// ctx.Data()

	co, err := NewCondition(filterData, core.LOGIC_ALL)
	if !ast.Nil(err) {
		t.Fatal(err)
	}
	// co, err := NewCondition([]interface{}{
	// 	[]interface{}{"has_magic_no_list", "has", "214"},
	// 	[]interface{}{"guide_step", "=", 12},
	// }, core.LOGIC_ALL)
	// if !ast.Nil(err) {
	// 	t.Fatal(err)
	// }
	// true
	ctx1 := core.WithData(ctx, map[string]interface{}{
		"has_magic_no_list":             []int32{214, 2},
		"is_finished_guide_step": true,
	})
	ast.True(co.Success(ctx1))

	// false
	ctx2 := core.WithData(ctx, map[string]interface{}{
		"has_magic_no_list":             []int32{214, 2},
		"is_finished_guide_step": false,
	})
	ast.False(co.Success(ctx2))

	// default
	ctx3 := core.WithData(ctx, map[string]interface{}{
		"has_magic_no_list": []int32{214, 2},
		// "is_finished_guide_step": true,
	})
	ast.False(co.Success(ctx3))
}
