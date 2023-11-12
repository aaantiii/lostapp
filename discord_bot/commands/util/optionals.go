package util

func OptionalString(s string) *string {
	return &s
}

func OptionalInt(v int) *int {
	return &v
}

func OptionalFloat(v float64) *float64 {
	return &v
}

func OptionalBool(v bool) *bool {
	return &v
}
