package istrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2Int64(t *testing.T) {
	ast := assert.New(t)
	type testCase struct {
		InputA   string
		Expected int64
	}

	cases := []testCase{
		{
			InputA:   "123",
			Expected: 123,
		},
		{
			InputA:   "asd",
			Expected: 0,
		},
		{
			InputA:   "-1",
			Expected: -1,
		},
		{
			InputA:   "0",
			Expected: 0,
		},
	}

	for _, c := range cases {
		ast.Equal(c.Expected, String2Int64(c.InputA))
	}
}
