package staticQueryBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/userData"
	"math/rand"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&rebornBot{10})
}

type rebornBot struct {
	priority int //[0~1000)
}

func (r rebornBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if req.IsAtMe && (req.Content == "转生") {
		return true, false
	}
	return false, true
}

func (r rebornBot) Process(req *plugin.Request) []*plugin.Result {
	user := userData.GetUser(req.Udid)
	if len(user.CardIndex) == len(cards.Cards)-1 {
		user.CardIndex = nil
		user.Static.RebornCount++
		volunter := rand.Intn(3)
		user.RebornEggNumber += volunter
		return []*plugin.Result{{Content: fmt.Sprintf("转生成功,转生次数%d,送转生券%d张", user.Static.RebornCount, volunter)}}
	} else {
		return []*plugin.Result{{Content: fmt.Sprintf("您满图鉴了么?,%s", user.GetCollection())}}
	}

}

func (r rebornBot) Priority() int {
	return r.priority
}
