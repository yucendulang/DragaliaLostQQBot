package mosaicBot

import (
	"fmt"
	"image"
	"image/draw"
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
const width = 70

var level = map[int]levelInfo{
	1: {width / 9, "初级", ""},
	2: {width / 8, "中级", ""},
	3: {width / 7, "高级", ""},
	4: {width / 6, "超级", ""},
	5: {width / 5, "入门", "真"},
	6: {width / 4, "中级", "真"},
	7: {width / 3, "高级", "真"},
	8: {width / 2, "超级", "真"},
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
	defer userData.SaveUserByUDID(req.Udid)
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
	//todo 第0个cards没东西 roll到就panic 将来重构掉
	card := cards.Cards[rand.Intn(len(cards.Cards)-1)+1]
	img := summon.GetCardImage(card.IconUrl)
	dest := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx()-10, img.Bounds().Dy()-10))

	draw.Draw(dest, dest.Rect, img, image.Point{X: 5, Y: 5}, draw.Over)
	lv, ok := level[user.MiniGame.Mosaic.Level]
	if !ok {
		user.MiniGame.Mosaic.Level = 1
		lv = level[1]
	}
	pic, _ := util.Mosaic(dest, lv.size-rand.Intn(1))
	user.MiniGame.Mosaic.Answer = card.Title
	user.MiniGame.Mosaic.StartTime = time.Now()
	return lv, pic
}

func (m *mosaicBot) Priority() int {
	return m.priority
}
