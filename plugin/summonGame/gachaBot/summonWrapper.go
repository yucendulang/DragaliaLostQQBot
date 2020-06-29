package gachaBot

import (
	"fmt"
	"image"
	"iotqq-plugins-demo/Go/achievement"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"sort"
)

type summonResult struct {
	image.Image
	string
}

func SummonALot(udid int64, num int, summonFunc func(*userData.User) summon.SummonRecord) []summonResult {
	defer userData.SaveUserByUDID(udid)
	user := userData.GetUser(udid)
	sr := summonALotGacha(user, num, summonFunc)
	//如果抽完卡满足了全图鉴
	if len(user.CardIndex) == len(cards.Cards)-1 {
		if user.Achieve(achievement.GaChaThemAll) {
			sr = append(sr, summonResult{nil, achievement.AchievementList[achievement.GaChaThemAll].Format("")})
		}
	}
	return sr
}

func summonALotGacha(user *userData.User, num int, summonFunc func(*userData.User) summon.SummonRecord) []summonResult {
	if user.SummonCardNum >= num {
		prevWater := user.Water
		res := summonFunc(user)
		changeWater := user.Water - prevWater
		user.SummonCardNum -= num
		if num == 10 {
			img := res.ImageFormatV2(user.SummonCardNum, user.Water)
			return []summonResult{{img, ""}}
		} else {
			res.StackCard()
			sort.Slice(res.Card, func(i, j int) bool {
				if res.Card[i].Star != res.Card[j].Star {
					return res.Card[i].Star > res.Card[j].Star
				}
				if res.Card[i].New != res.Card[j].New {
					return res.Card[i].New
				}
				return res.Card[i].StackNum > res.Card[j].StackNum
			})
			var sr []summonResult
			sSRCount := res.CountSSR()
			if num == 100 {
				achies := []int{achievement.SummonGreatThan30SSR, achievement.SummonGreatThan20SSR,
					achievement.SummonGreatThan10SSR, achievement.SummonEqual0SSR}
				for _, item := range achies {
					if achievement.AchievementList[item].Trigger(sSRCount) {
						if user.Achieve(item) {
							sr = append(sr, summonResult{nil, achievement.AchievementList[item].Format("")})
						}
						break
					}
				}
			}

			firstFlag := true
			for {
				OutStr := ""
				if firstFlag {
					OutStr += fmt.Sprintf("\n抽%d次,五星%d个,水滴:%d", num, sSRCount, changeWater)
				} else {
					OutStr += "好多new" + random.RandomGetSuffix()
				}

				//if len(res.Card) < 10 {
				//	sr = append(sr, summonResult{nil, "需要展示卡牌小于10张了,这个我暂时还不会"})
				//	break
				//}
				//if res.Card[10].Star == 5 {
				//	//if res.Card[10].New {
				//	OutStr += "\n命运之子啊~你还有更多的五星~让我慢慢展示给你"
				//}
				img := res.ImageFormatV2(user.SummonCardNum, user.Water)
				sr = append(sr, summonResult{img, OutStr})
				if len(res.Card) > 10 && res.Card[10].New {
					res.Card = res.Card[10:]
				} else {
					break
				}
			}
			return sr
		}
	} else {
		return []summonResult{{nil, fmt.Sprintf("您只有%d🎟,想抽%d,那召唤券就不够了%s", user.SummonCardNum, num, random.RandomGetSuffix())}}
	}
}
