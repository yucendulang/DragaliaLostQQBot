package gachaBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&gachaBot{8})
}

type gachaBot struct {
	priority int //[0~1000)
}

func (g *gachaBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if util.KeyWordTrigger(req.Content, "抽卡") || util.KeyWordTrigger(req.Content, "单抽") {
		return true, false
	}
	return false, true
}

func (g *gachaBot) Process(req *plugin.Request) *plugin.Result {

	user := userData.GetUser(req.Udid)
	if user.SummonCardNum >= 1 {
		res := summon.OneSummon(user)
		user.SummonCardNum--
		userData.UserDataSave()
		img := res.ImageFormatV2(user.SummonCardNum, user.Water)
		return &plugin.Result{
			Pic: img,
		}
	} else {
		return &plugin.Result{
			Content: fmt.Sprintf("%s召唤券不够了%s", req.NickName, random.RandomGetSuffix()),
		}
	}
}

func (g *gachaBot) Priority() int {
	return g.priority
}
