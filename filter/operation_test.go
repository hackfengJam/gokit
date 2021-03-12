package filter

import (
	"testing"

	"github.com/hackfengJam/filter/core"
	"github.com/stretchr/testify/assert"
)

type opTestCase struct {
	input         []interface{}
	expected      bool
	expectedError bool
}

func testCases(ctx *core.Context, ast *assert.Assertions, tests []opTestCase) {
	for _, c := range tests {
		op := GetOperation(c.input[1].(string))
		ast.NotNil(op)
		v := CreateVariable(c.input[0].(string))
		ast.NotNil(v)
		pv, err := op.PrepareValue(c.input[2])
		if c.expectedError {
			ast.Error(err)
			continue
		}
		ast.NoError(err)
		ast.Equal(c.expected, op.Run(ctx, v, pv), c.input)
	}
}

func setupTest() *core.Context {
	ctx := core.NewContext()
	return ctx
}

func TestGetOperation(t *testing.T) {
	ast := assert.New(t)
	ctx := setupTest()

	data := map[string]interface{}{
		"area": map[string]interface{}{
			"zipcode": 200211,
			"city":    "shanghai",
		},
		"birthday":  "1985-11-21",
		"height":    "178",
		"age":       25,
		"fav_books": "book1,book2,book3",
		"pets":      []interface{}{"dog", "cat"},
	}
	ctx = core.WithData(ctx, data)

	tests := []opTestCase{
		{[]interface{}{"data.fav_books", "has", "book1,book3"}, true, false},
		{[]interface{}{"data.fav_books", "has", "book1,book3,book4"}, false, false},
		{[]interface{}{"data.pets", "any", "cat,pig"}, true, false},
		{[]interface{}{"data.pets", "any", "rabbit,fly"}, false, false},
		{[]interface{}{"data.pets", "none", []interface{}{"pig", "rabbit"}}, true, false},
		{[]interface{}{"data.pets", "none", []interface{}{"dog", "rabbit"}}, false, false},
	}
	testCases(ctx, ast, tests)
}
