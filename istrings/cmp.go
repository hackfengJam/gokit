package istrings

import (
	"sort"

	"github.com/google/go-cmp/cmp"
)

// StringSetEqual String Equal
func StringSetEqual(a []string, b []string) bool {
	lnA := len(a)
	lnB := len(b)
	if lnA != lnB {
		return false
	}

	opt := cmp.Comparer(func(a []string, b []string) bool {
		if lnA != lnB {
			return false
		}
		iA := make([]string, lnA)
		copy(iA, a)
		iB := make([]string, lnB)
		copy(iB, b)

		sort.Slice(iA, func(i, j int) bool {
			return iA[i] < iA[j]
		})
		sort.Slice(iB, func(i, j int) bool {
			return iB[i] < iB[j]
		})
		return cmp.Equal(iA, iB)
	})

	return cmp.Equal(a, b, opt)
}
