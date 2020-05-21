package queryBot

import (
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/userData"
	"strings"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&queryBot{6})
}

type queryBot struct {
	priority int //[0~1000)
}

func (q *queryBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if req.Content == "圣城" {
		return true, false
	}
	return false, true
}

func (q *queryBot) Process(req *plugin.Request) *plugin.Result {
	user := userData.GetUser(req.Udid)
	var Outer = []string{user.GetMyHitRate(), user.GetAccountInfo(), user.GetCollection(), user.GetBuildInfo()}

	return &plugin.Result{Content: strings.Join(Outer, "\n")}
}

func (q *queryBot) Priority() int {
	return q.priority
}
