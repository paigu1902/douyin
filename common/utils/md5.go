package utils

import (
	"crypto/md5"
	"fmt"
)

func CalMD5(message string) string {
	data := []byte(message)
	return fmt.Sprintf("%x", md5.Sum(data))
}
