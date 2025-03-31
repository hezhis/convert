package convert

import "strconv"

func ConvertAs[T Integer](s string) (T, error) {
	var value T
	var i int64
	var err error

	switch any(value).(type) {
	case int, int8, int16, int32, int64:
		i, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, err
		}
		value = T(i)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		u, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return 0, err
		}
		value = T(u)
	}

	return value, nil
}

func Convert[T Integer](val T) string {
	switch any(val).(type) {
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(int64(val), 10)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return strconv.FormatUint(uint64(val), 10)
	}

	return ""
}
