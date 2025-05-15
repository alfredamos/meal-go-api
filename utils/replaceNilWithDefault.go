package utils


func ReplaceNilWithDefaultInt(ptr *int, defaultValue int) *int {
	if ptr == nil {
		ptr = new(int)
		*ptr = defaultValue
	}

	return ptr
}

func ReplaceNilWithDefaultFloat(ptr *float64, defaultValue float64) *float64 {
	if ptr == nil {
		ptr = new(float64)
		*ptr = defaultValue
	}

	return ptr
}

func ReplaceNilWithDefaultStr(ptr *string, defaultValue string) *string {
	if ptr == nil || *ptr == "" {
		ptr = new(string)
		*ptr = defaultValue
	}

	return ptr
}

