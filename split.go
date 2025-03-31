package convert

import (
	"strconv"
	"strings"
)

// Integer 是一个类型约束接口，限定只能支持 int8, int16, int32, int64
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// SplitAsSlice 将字符串按分隔符切割成整数切片
func SplitAsSlice[T Integer](s, sep string) ([]T, error) {
	strSlice := strings.Split(s, sep)
	var result []T

	var value T

	switch any(value).(type) {
	case int, int8, int16, int32, int64:
		for _, fs := range strSlice {
			if i, err := strconv.ParseInt(fs, 10, 64); err != nil {
				return nil, err
			} else {
				result = append(result, T(i))
			}
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		for _, fs := range strSlice {
			if i, err := strconv.ParseUint(fs, 10, 64); err != nil {
				return nil, err
			} else {
				result = append(result, T(i))
			}
		}
	}

	return result, nil
}
