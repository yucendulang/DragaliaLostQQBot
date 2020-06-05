package staticQueryBot

import (
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/userData"
	"strings"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&staticQueryBot{9})
}

type staticQueryBot struct {
	priority int //[0~1000)
}

func (s *staticQueryBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if req.IsAtMe && (req.Content == "统计" || req.Content == "成就") {
		return true, false
	}
	return false, true
}

func (s *staticQueryBot) Process(req *plugin.Request) []*plugin.Result {
	user := userData.GetUser(req.Udid)
	var Outer = []string{user.GetStatic(), user.GetAchievement()}
	return []*plugin.Result{{Content: strings.Join(Outer, "\n")}}
}

func (s *staticQueryBot) Priority() int {
	return s.priority
}
