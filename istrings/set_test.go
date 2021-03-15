package istrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringList2Set(t *testing.T) {
	ast := assert.New(t)
	type testCase struct {
		InputA   []string
		Expected []string
	}

	cases := []testCase{
		{
			InputA:   []string{"a", "b", "b", "b", "c"},
			Expected: []string{"a", "b", "c"},
		},
		{
			InputA:   []string{"a", "b", "a", "b", "a", "b", "a", "b", "a", "b", "c"},
			Expected: []string{"a", "b", "c"},
		},
		{
			InputA:   []string{"a", "a", "a", "a", "a", "b", "c"},
			Expected: []string{"a", "b", "c"},
		},
		{
			InputA:   []string{"a", "b", "c", "b", "c", "b", "c", "b", "c"},
			Expected: []string{"a", "b", "c"},
		},
		{
			InputA:   []string{"a", "c", "b"},
			Expected: []string{"a", "b", "c"},
		},
	}

	for _, c := range cases {
		ast.True(StringSetEqual(c.Expected, StringList2Set(c.InputA)))
	}
}
