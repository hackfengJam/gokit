package istrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSetEqual(t *testing.T) {
	ast := assert.New(t)
	type testCase struct {
		InputA   []string
		InputB   []string
		Expected bool
	}

	cases := []testCase{
		{
			InputA:   []string{"a", "c", "b"},
			InputB:   []string{"a", "b", "c"},
			Expected: true,
		},
		{
			InputA:   []string{"a", "b", "c"},
			InputB:   []string{"a", "b", "c"},
			Expected: true,
		},
		{
			InputA:   []string{"c", "b", "a"},
			InputB:   []string{"a", "b", "c"},
			Expected: true,
		},
		{
			InputA:   []string{"a", "b", "b", "b", "c"},
			InputB:   []string{"a", "b", "c"},
			Expected: false,
		},
	}

	for _, c := range cases {
		ast.Equal(c.Expected, StringSetEqual(c.InputA, c.InputB))
	}
}
