package utils

func Response(status int, data interface{}, err error) (result map[string]interface{}) {
	result = map[string]interface{}{
		"status": status,
		"data":   data,
		"error":  err.Error(),
	}

	return
}
