package cards

import (
	"iotqq-plugins-demo/Go/util"
	"math/rand"
)

type cardManager []*CardCollectionV2

var GachaPoolCardMgr cardManager
var NotGachaPoolCardMgr cardManager

func (c cardManager) PickUpOne() *CardCollectionV2 {
	ran1 := rand.Intn(4)
	if ran1 == 0 {
		return c[0]
	} else if ran1 == 1 {
		return c[1]
	} else {
		ran := rand.Intn(len(c))
		return c[ran]
	}
}

func (c cardManager) PickUp(index int) *CardCollectionV2 {
	if index >= len(c) {
		index = 0
	}
	return c[index]
}

func init() {
	collectConfigs := initConfig()
	initCards()
	for _, config := range collectConfigs {
		var cardCollection CardCollectionV2
		var pickUpCards []int
		for _, poolConfig := range config.configs {
			var cards []Card
			for _, index := range poolConfig.pickUpCards {
				cards = append(cards, Cards[index])
			}
			cardCollection.cardSets = append(cardCollection.cardSets, &CardSet{
				star:     poolConfig.star,
				cardType: poolConfig.cardType,
				rareType: poolConfig.rareType,
				Prob:     poolConfig.Prob,
				cards:    cards,
			})
			cardCollection.TopBannerUrl = config.topBannerUrl
			cardCollection.ProbFix = config.probFix
			pickUpCards = append(pickUpCards, poolConfig.pickUpCards...)
		}

		for _, card := range Cards {
			if util.IntContain(card.ID, pickUpCards) {
				continue
			}
			for _, cardSet := range cardCollection.cardSets {
				if cardSet.star == card.Star && util.IntContain(card.RareType, cardSet.rareType) && cardSet.cardType == card.CardType {
					cardSet.cards = append(cardSet.cards, card)
					break
				}
			}
		}
		GachaPoolCardMgr = append(GachaPoolCardMgr, &cardCollection)
	}

	//初始化非扭蛋卡池
	{
		var cardCollection CardCollectionV2
		cardCollection.cardSets = append(cardCollection.cardSets, &CardSet{
			star:     0,
			cardType: 0,
			rareType: []int{RareTypeStory, RareTypeEvent},
			Prob:     0,
			cards:    []Card{},
		})
		for _, card := range Cards {
			if util.IntContain(card.RareType, cardCollection.cardSets[0].rareType) {
				cardCollection.cardSets[0].cards = append(cardCollection.cardSets[0].cards, card)
			}
		}
		NotGachaPoolCardMgr = append(NotGachaPoolCardMgr, &cardCollection)
	}
}
