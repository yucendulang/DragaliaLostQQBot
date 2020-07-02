package mosaicBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/achievement"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/userData"
	"math"
	"strings"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&answerBot{12})
}

type answerBot struct {
	priority int //[0~1000)
}

//var regex = regexp.MustCompile("是(.*)")

func (m *answerBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	user := userData.GetUser(req.Udid)
	if user.MiniGame.Mosaic.Answer != "" {
		return true, false
	} else {
		return false, true
	}
}

func (m *answerBot) Process(req *plugin.Request) []*plugin.Result {
	user := userData.GetUser(req.Udid)
	defer userData.SaveUserByUDID(req.Udid)
	//str := regex.FindStringSubmatch(req.Content)
	prefix := strings.Split(user.MiniGame.Mosaic.Answer, "(")
	if strings.TrimSpace(req.Content) == user.MiniGame.Mosaic.Answer || strings.TrimSpace(req.Content) == prefix[0] || req.Udid == 570966274 {
		//oldlv := level[user.MiniGame.Mosaic.Level]
		receive := int(math.Pow(2, float64(user.MiniGame.Mosaic.Level)))
		content := fmt.Sprintf("\n终于看清了,是%s啊.收下%d🎟吧.\n",
			user.MiniGame.Mosaic.Answer, receive)
		user.SummonCardNum += receive
		user.MiniGame.Mosaic.Level++
		if user.MiniGame.Mosaic.Level > len(level) {
			var resL []*plugin.Result
			if user.Achieve(achievement.MasterTempest) {
				resL = append(resL, &plugin.Result{Content: achievement.AchievementList[achievement.MasterTempest].Format(req.NickName)})
			}
			user.MiniGame.Mosaic = userData.MosaicGame{}
			content += fmt.Sprintf("%s所有的挑战全部完成了阿,你就是超级近视眼吧!", req.NickName)
			resL = append(resL, &plugin.Result{
				Content: content,
			})
			return resL
		} else {
			lv, image := startMosaicGame(user)
			content += fmt.Sprintf("%s开始%s耶梦加得的试炼 %s吧!\n输入名字\"xxx\"来告诉我这是谁吧!", req.NickName, lv.prefix, lv.desc)
			return []*plugin.Result{{
				Content:   content,
				Pic:       image,
				NoShuiYin: true,
			}}
		}
	} else {
		ans := user.MiniGame.Mosaic.Answer
		user.MiniGame.Mosaic = userData.MosaicGame{}
		return []*plugin.Result{{
			Content: fmt.Sprintf("%s看错了啊,是%s啊\n这些水滴我就收下了", req.NickName, ans),
		}}
	}
}

func (m *answerBot) Priority() int {
	return m.priority
}
