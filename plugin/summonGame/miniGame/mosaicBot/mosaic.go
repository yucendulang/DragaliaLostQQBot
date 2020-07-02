package mosaicBot

import (
	"fmt"
	"image"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"math/rand"
	"time"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&mosaicBot{11})
}

const keyWord = "耶梦加得的试炼"

var level = map[int]levelInfo{
	1: {5, "初级", ""},
	2: {8, "中级", ""},
	3: {10, "高级", ""},
	4: {13, "超级", ""},
	5: {16, "入门", "真"},
	6: {20, "中级", "真"},
	7: {26, "高级", "真"},
	8: {40, "超级", "真"},
}

type levelInfo struct {
	size   int
	desc   string
	prefix string
}

type mosaicBot struct {
	priority int //[0~1000)
}

func (m *mosaicBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	return plugin.NewCommonPrefixTriggerFunc(keyWord)(req)
}

func (m *mosaicBot) Process(req *plugin.Request) []*plugin.Result {
	user := userData.GetUser(req.Udid)
	//判断是否能开启游戏
	water := 100000
	if user.Water < water {
		return []*plugin.Result{{Content: "需要10w💧参加耶梦加得的试炼" + random.RandomGetSuffix()}}
	}

	if user.MiniGame.Mosaic.Level != 0 {
		return []*plugin.Result{{Content: "好像你正在跟耶耶子玩耍" + random.RandomGetSuffix()}}
	}
	if time.Now().Sub(user.MiniGame.Mosaic.StartTime) > time.Minute*10 {
		user.MiniGame.Mosaic = userData.MosaicGame{}
	}
	lv, pic := startMosaicGame(user)
	user.Water -= water
	return []*plugin.Result{{
		Content:   fmt.Sprintf("\n这个东西带上好晕啊,让我康康这是谁-%s耶梦加得的试炼 %s\n输入名字\"xxx\"%s来告诉我这是谁吧!", lv.prefix, lv.desc, req.NickName),
		Pic:       pic,
		NoShuiYin: true,
	}}
}

func startMosaicGame(user *userData.User) (levelInfo, image.Image) {
	card := cards.Cards[rand.Intn(len(cards.Cards))]
	img := summon.GetCardImage(card.IconUrl)
	lv, ok := level[user.MiniGame.Mosaic.Level]
	if !ok {
		user.MiniGame.Mosaic.Level = 1
		lv = level[1]
	}
	pic, _ := util.Mosaic(img, lv.size)
	user.MiniGame.Mosaic.Answer = card.Title
	user.MiniGame.Mosaic.StartTime = time.Now()
	return lv, pic
}

func (m *mosaicBot) Priority() int {
	return m.priority
}
