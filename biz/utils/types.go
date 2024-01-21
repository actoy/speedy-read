package utils

import "strconv"

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringToInt64List(sList []string) []int64 {
	result := make([]int64, 0)
	for _, s := range sList {
		result = append(result, StringToInt64(s))
	}
	return result
}

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
