package plugin

import "strings"

func NewCommonPrefixTriggerFunc(prefix string) func(request *Request) (bool, bool) {
	return func(req *Request) (bool, bool) {
		if req.IsAtMe && strings.HasPrefix(req.Content, prefix) {
			return true, false
		}
		return false, true
	}
}
