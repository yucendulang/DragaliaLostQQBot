package repeatBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/building"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&repeatBot{priority: 999})
}

type repeatBot struct {
	priority int //[0~1000)
}

func (r repeatBot) Priority() int {
	return r.priority
}

func (r repeatBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	//if req.Udid==570966274{
	//	return true,true
	//}
	content := util.FixSentense(req.Content)
	if len(content) <= 30 && len(content) > 0 {
		//fmt.Println("enter repeatbot trigger")
		user := userData.GetUser(req.Udid)
		eff := building.GetBuildEffect(user.BuildIndex)
		res := summon.OneSummon(&userData.User{UnHitNumber: eff.RepeatProbability / 2})
		if res.Card[0].Star == 5 {
			return true, false
		}
	}
	return false, true
}

func (r repeatBot) Process(req *plugin.Request) *plugin.Result {
	res := &plugin.Result{}
	user := userData.GetUser(req.Udid)
	eff := building.GetBuildEffect(user.BuildIndex)
	num := int(float32(RandomSummonCard()*10) * eff.GetExtraRepeatBonus())
	user.SummonCardNum += num
	res.Content = fmt.Sprintf("%s%s\n(送%s殿下%d张🎟", util.FixSentense(req.Content), random.RandomGetSuffix(), req.NickName, num)
	return res
}
