package repeatBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/achievement"
	"iotqq-plugins-demo/Go/building"
	"iotqq-plugins-demo/Go/common"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"time"
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
	//if req.Udid == 570966274 {
	//	return true, true
	//}
	content := util.FixSentense(req.Content)
	if len(content) <= 30 && len(content) > 0 {
		//fmt.Println("enter repeatbot trigger")
		user := userData.GetUser(req.Udid)
		eff := building.GetBuildEffect(user.BuildIndex)
		unHitNumber := eff.RepeatProbability/2 - user.Static.VolunterReiceiveTime/10 - user.Static.VRTPeriod*5
		fmt.Println("unHitNumber:", unHitNumber)
		res := summon.OneSummon(&userData.User{UnHitNumber: unHitNumber})
		if res.Card[0].Star == 5 {
			return true, false
		}
	}
	return false, true
}

func (r repeatBot) Process(req *plugin.Request) []*plugin.Result {
	defer userData.SaveUserByUDID(req.Udid)
	res := &plugin.Result{}
	resL := []*plugin.Result{res}
	user := userData.GetUser(req.Udid)
	eff := building.GetBuildEffect(user.BuildIndex)
	base := RandomSummonCard() * 10
	num := int(float32(base) * eff.GetExtraRepeatBonus())
	if num > user.Static.VolunterReiceiveMax {
		user.Static.VolunterReiceiveMax = num
	}
	if achievement.AchievementList[achievement.ReiceiveLotVolunter].Trigger(num) {
		if user.Achieve(achievement.ReiceiveLotVolunter) {
			resL = append(resL, &plugin.Result{Content: achievement.AchievementList[achievement.ReiceiveLotVolunter].Format(req.NickName)})
		}
	}
	user.Static.VolunterReiceiveTime++
	user.Static.VRTPeriod++
	user.SummonCardNum += num

	if base == 10 && user.LastVolunterGetTime.Add(common.VolunterMineProductPeriod).Sub(time.Now()).Minutes() < 30 {
		if user.Achieve(achievement.CoinMineRefresh) {
			resL = append(resL, &plugin.Result{Content: achievement.AchievementList[achievement.CoinMineRefresh].Format(req.NickName)})
		}
	}
	user.LastVolunterGetTime = time.Now()
	res.Content = fmt.Sprintf("%s%s\n(送%s殿下%d张🎟", util.FixSentense(req.Content), random.RandomGetSuffix(), req.NickName, num)

	return resL
}
