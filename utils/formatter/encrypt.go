package formatter

import (
	"crypto/md5"
	"fmt"
)

func EncryptMd5(word string) string {
	h := md5.New()
	return fmt.Sprintf("%x", h.Sum([]byte(word)))
}
