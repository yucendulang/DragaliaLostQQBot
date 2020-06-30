package staticQueryBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/common"
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
	gachaNum, _ := cards.GetGachaCardsNum(user.CardIndex)
	if gachaNum == common.GachaPoolNum {
		var cardsIndexNew []int
		for _, index := range user.CardIndex {
			if cards.Cards[index].RareType == cards.RareTypeStory || cards.Cards[index].RareType == cards.RareTypeEvent {
				cardsIndexNew = append(cardsIndexNew, index)
			}
		}
		user.CardIndex = cardsIndexNew
		user.Static.RebornCount++
		volunter := rand.Intn(3)
		user.RebornEggNumber += volunter
		userData.SaveUserByUDID(user.Udid)
		return []*plugin.Result{{Content: fmt.Sprintf("转生成功,转生次数%d,送转生券%d张", user.Static.RebornCount, volunter)}}
	} else {
		return []*plugin.Result{{Content: fmt.Sprintf("您蛋池满图鉴了么?,蛋池收集进度%d/%d", gachaNum, common.GachaPoolNum)}}
	}

}

func (r rebornBot) Priority() int {
	return r.priority
}
