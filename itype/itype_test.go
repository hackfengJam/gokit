package itype

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	tests := []struct {
		obj   interface{}
		value float64
	}{
		{
			obj:   false,
			value: 0.0,
		},
		{
			obj:   true,
			value: 1.0,
		},
		{
			obj:   int32(3),
			value: 3,
		},
		{
			obj:   float32(1.5),
			value: 1.5,
		},
		{
			obj:   "2.345",
			value: 2.345,
		},
		{
			obj:   "",
			value: 0.0,
		},
		{
			obj:   nil,
			value: 0.0,
		},
		{
			obj:   "invalid",
			value: 0.0,
		},
		{
			obj:   struct{}{},
			value: 0.0,
		},
	}

	for _, test := range tests {
		f := Float(test.obj)
		if math.Abs(f-float64(test.value)) > 1e-9 {
			t.Errorf("%v != %f", test.obj, test.value)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		obj   interface{}
		value int64
	}{
		{
			obj:   false,
			value: 0,
		},
		{
			obj:   true,
			value: 1,
		},
		{
			obj:   int32(3),
			value: 3,
		},
		{
			obj:   "2",
			value: 2,
		},
		{
			obj:   "",
			value: 0,
		},
		{
			obj:   "-1",
			value: -1,
		},
		{
			obj:   nil,
			value: 0,
		},
		{
			obj:   struct{}{},
			value: 0,
		},
	}

	for _, test := range tests {
		f := Int(test.obj)
		if f != test.value {
			t.Errorf("%v !=> %d", test.obj, test.value)
		}
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		obj   interface{}
		value uint64
	}{
		{
			obj:   false,
			value: 0,
		},
		{
			obj:   true,
			value: 1,
		},
		{
			obj:   int32(3),
			value: 3,
		},
		{
			obj:   "2",
			value: 2,
		},
		{
			obj:   "",
			value: 0,
		},
		{
			obj:   "-1",
			value: 0,
		},
		{
			obj:   nil,
			value: 0,
		},
		{
			obj:   struct{}{},
			value: 0,
		},
	}

	for _, test := range tests {
		f := Uint(test.obj)
		if f != test.value {
			t.Errorf("%v != %d. %d", test.obj, test.value, f)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		obj   interface{}
		value string
	}{
		{
			obj:   false,
			value: "",
		},
		{
			obj:   true,
			value: "1",
		},
		{
			obj:   3,
			value: "3",
		},
		{
			obj:   2.1,
			value: "2.1",
		},
		{
			obj:   nil,
			value: "",
		},
		{
			obj:   struct{}{},
			value: "",
		},
	}

	for _, test := range tests {
		f := String(test.obj)
		if f != test.value {
			t.Errorf("%v != %s. %s", test.obj, test.value, f)
		}
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		obj   interface{}
		value bool
	}{
		{
			obj:   false,
			value: false,
		},
		{
			obj:   true,
			value: true,
		},
		{
			obj:   3,
			value: true,
		},
		{
			obj:   0.0,
			value: false,
		},
		{
			obj:   nil,
			value: false,
		},
		{
			obj:   struct{}{},
			value: true,
		},
		{
			obj:   "",
			value: false,
		},
		{
			obj:   "false",
			value: false,
		},
		{
			obj:   "off",
			value: false,
		},
		{
			obj:   "no",
			value: false,
		},
		{
			obj:   "0",
			value: false,
		},
		{
			obj:   "1",
			value: true,
		},
		{
			obj:   struct{}{},
			value: true,
		},
	}

	for _, test := range tests {
		f := Bool(test.obj)
		if f != test.value {
			t.Errorf("%+v != %v. %v", test.obj, test.value, f)
		}
	}
}

func TestType(t *testing.T) {
	tests := []struct {
		in    interface{}
		check Type
	}{
		{1, NUMBER},
		{1.2, NUMBER},
		{false, BOOL},
		{"1", STRING},
		{[]int{1}, ARRAY},
		{[1]int{1}, ARRAY},
		{map[int]int{1: 1}, MAP},
		{nil, NULL},
		{func() {}, UNKNOWN},
	}

	for i, test := range tests {
		assert.Equal(t, test.check, GetType(test.in), "case %d: %v", i, test.in)
	}
}

func BenchmarkIntFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(3)
	}
}

func BenchmarkIntSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int("3")
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool("false")
	}
}
