package stickerBot

import (
	"iotqq-plugins-demo/Go/achievement"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"math/rand"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&stickerBot{2})
}

type stickerBot struct {
	priority int //[0~1000)
}

func (s *stickerBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	content := util.FixSentense(req.Content)
	url, b := IsStickerKey(content)
	req.ExtraInfo = url
	if b {
		return true, false
	} else {
		return false, true
	}
}

func (s *stickerBot) Process(req *plugin.Request) []*plugin.Result {
	res := &plugin.Result{PicUrl: req.ExtraInfo.(string)}
	resL := []*plugin.Result{res}
	if rand.Intn(100) < 5 {
		url2, _ := IsStickerKey("磕头")
		res.PicUrl = url2
		res.Content = "\n轨迹阵亡...放弃 再次挑战"
		if userData.GetUser(req.Udid).Achieve(achievement.StickerFailed) {
			resL = append(resL, &plugin.Result{Content: achievement.AchievementList[achievement.StickerFailed].Format(req.NickName)})
		}
	}
	return resL
}

func (s *stickerBot) Priority() int {
	return s.priority
}
