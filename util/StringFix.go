package util

import "strings"

var nameSuffix = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", " ", "_", "sw", "—", "SW"}
var sentenseSuffix = []string{"?", "!", ".", "呢", "嘛", "吗", "呐", "！", "？", "了"}

func FixName(name string) string {
	return DelStringTail(name, nameSuffix)
}

func FixSentense(sentense string) string {
	return DelStringTail(sentense, sentenseSuffix)
}

func DelStringTail(str string, suffix []string) string {
	for _, s := range suffix {
		res := strings.TrimSuffix(str, s)
		if len(res) != len(str) {
			return DelStringTail(res, suffix)
		}
	}
	return str
}
