package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// Response just return response info
func Response(status int, data interface{}, err error) (result map[string]interface{}) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	result = map[string]interface{}{
		"status": status,
		"data":   data,
		"error":  errMsg,
	}

	return
}

// Md5Encode will encode string with md5 method
func Md5Encode(resource string) string {
	h := md5.New()
	io.WriteString(h, resource)
	return string(h.Sum(nil))
}

func Sha256Encode(resource string) string {
	result := sha256.Sum256([]byte(resource))
	return hex.EncodeToString(result[:])
}
