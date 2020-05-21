package util

import "strings"

func KeyWordTrigger(content, key string) bool {
	return strings.HasPrefix(content, key) || strings.HasSuffix(content, key)
}
