package util

func GetPageOffset(page int, pageSize int) int {
	result := 0

	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
