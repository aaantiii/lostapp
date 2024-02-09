package util

func IntPtr(v int) *int {
	return &v
}

func FloatPtr(v float64) *float64 {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}
