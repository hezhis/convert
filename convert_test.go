package convert

import "testing"

func TestConvertAs(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-456", -456},
		{"789", 789},
	}

	for _, test := range tests {
		result, err := ConvertAs[int](test.input)
		if err != nil {
			t.Errorf("ConvertAs(%s) returned error: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("ConvertAs(%s) = %d; want %d", test.input, result, test.expected)
		}
	}

	// 测试 uint 类型
	testsUint := []struct {
		input    string
		expected uint
	}{
		{"123", 123},
		{"456", 456},
		{"789", 789},
	}

	for _, test := range testsUint {
		result, err := ConvertAs[uint](test.input)
		if err != nil {
			t.Errorf("ConvertAs(%s) returned error: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("ConvertAs(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}
