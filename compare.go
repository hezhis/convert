package convert

// Ordered 约束接口，要求实现 Less 方法
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Max 返回两个数中较大的那个
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min 返回两个数中较小的那个
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
