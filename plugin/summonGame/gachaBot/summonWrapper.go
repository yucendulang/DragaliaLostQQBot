package gachaBot

import (
	"fmt"
	"image"
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
	defer userData.UserDataSave()
	user := userData.GetUser(udid)
	if user.SummonCardNum >= num {
		res := summonFunc(user)
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
			for {
				OutStr := ""
				if res.Card[10].Star == 5 {
					if res.Card[10].New {
						OutStr += "\n命运之子啊~你还有更多的五星~让我慢慢展示给你"
					} else {
						ssrNum := 0
						for i := 10; i < len(res.Card); i++ {
							if res.Card[i].Star == 5 {
								ssrNum++
							} else {
								break
							}
						}
						OutStr += fmt.Sprintf("\n没有更多的new了,未展示的虹共计%d个", ssrNum)
					}
				}
				img := res.ImageFormatV2(user.SummonCardNum, user.Water)
				sr = append(sr, summonResult{img, OutStr})
				if res.Card[10].Star == 5 && res.Card[10].New {
					res.Card = res.Card[10:]
				} else {
					break
				}
			}
			return sr
		}
	} else {
		return []summonResult{{nil, "召唤券不够了" + random.RandomGetSuffix()}}
	}
}
