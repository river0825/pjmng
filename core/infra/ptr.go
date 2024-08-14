package infra

func Ptr[T any](v T) *T {
	return &v
}

func Val[T any](v *T, defaultVal T) T {
	if v != nil {
		return *v
	}

	return defaultVal
}
