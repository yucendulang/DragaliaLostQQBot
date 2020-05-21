package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func HashURL(url string) string {
	h := md5.New()
	io.WriteString(h, url)
	return fmt.Sprintf("%x", h.Sum(nil))
}
