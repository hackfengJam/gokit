package istrings

import "strconv"

// String2Int64 string转int64
func String2Int64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
