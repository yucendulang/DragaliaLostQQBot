package cards

type cardMgrConfig struct {
	configs      []*CardSetConfig
	topBannerUrl string
}

type CardSetConfig struct {
	star, cardType, rareType, Prob int
	pickUpCards                    []int
}

func initConfig() []*cardMgrConfig {
	var res []*cardMgrConfig
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 150, pickUpCards: []int{228, 229, 230}},
		{star: 5, Prob: 80, pickUpCards: []int{231}},
		{star: 5, cardType: 1, rareType: 3, Prob: 3 * 8},
		{star: 5, cardType: 2, rareType: 3, Prob: 5},
		{star: 5, cardType: 1, rareType: 1, Prob: 146},
		{star: 5, cardType: 2, rareType: 1, Prob: 195},
		{star: 5, Prob: 700, pickUpCards: []int{241}},
		{star: 4, cardType: 1, rareType: 1, Prob: 505},
		{star: 4, cardType: 2, rareType: 1, Prob: 395},
		{star: 3, cardType: 1, rareType: 1, Prob: 4700},
		{star: 3, cardType: 2, rareType: 1, Prob: 3100},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/6/66/Gala_Dragalia_%28May_2020%29_Summon_Top_Banner.png?version=a9d11303eea27a11e68dcc78bce6012b"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 80, pickUpCards: []int{194}},
		{star: 5, cardType: 1, rareType: 1, Prob: 120},
		{star: 5, cardType: 2, rareType: 1, Prob: 200},
		{star: 4, cardType: 1, rareType: 1, Prob: 800},
		{star: 4, cardType: 2, rareType: 1, Prob: 800},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/8/8c/Fire_Emblem_Kindred_Ties_%28Part_Two%29_Summon_Top_Banner.png?version=f56e65d83f19414f5bfb4311cd7e65b5"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 100, pickUpCards: []int{195, 196}},
		{star: 5, cardType: 1, rareType: 1, Prob: 120},
		{star: 5, cardType: 2, rareType: 1, Prob: 180},
		{star: 4, cardType: 1, rareType: 1, Prob: 800},
		{star: 4, cardType: 2, rareType: 1, Prob: 800},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/d/db/Fire_Emblem_Kindred_Ties_%28Part_One%29_Summon_Top_Banner.png?version=5432df699926a047193a1d2783a32b7d"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 150, pickUpCards: []int{197, 198, 199}},
		{star: 5, cardType: 1, rareType: 1, Prob: 120},
		{star: 5, cardType: 2, rareType: 1, Prob: 180},
		{star: 4, cardType: 1, rareType: 1, Prob: 800},
		{star: 4, cardType: 2, rareType: 1, Prob: 800},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/4/4b/Fire_Emblem_Lost_Heroes_%28Summon_Showcase%29_Summon_Top_Banner.png?version=d21f899d16f5d5d4992c53360b6048af"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 250, pickUpCards: []int{200, 201, 11, 149, 164}},
		{star: 5, cardType: 1, rareType: 1, Prob: 50},
		{star: 5, cardType: 2, rareType: 1, Prob: 100},
		{star: 4, Prob: 400, pickUpCards: []int{88, 57}},
		{star: 4, cardType: 1, rareType: 1, Prob: 600},
		{star: 4, cardType: 2, rareType: 1, Prob: 600},
		{star: 3, Prob: 1000, pickUpCards: []int{93}},
		{star: 3, cardType: 1, rareType: 1, Prob: 4300},
		{star: 3, cardType: 2, rareType: 1, Prob: 2700},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/c/c5/Dulcet_Delights_Summon_Top_Banner.png?version=320deb12822bd305bb8979cc214a8a2b"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 200, pickUpCards: []int{202, 203, 10, 138}},
		{star: 5, cardType: 1, rareType: 1, Prob: 80},
		{star: 5, cardType: 2, rareType: 1, Prob: 120},
		{star: 4, Prob: 600, pickUpCards: []int{67, 52, 169}},
		{star: 4, cardType: 1, rareType: 1, Prob: 500},
		{star: 4, cardType: 2, rareType: 1, Prob: 500},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/4/45/Monster_Hunter_Primal_Crisis_%28Part_Two%29_Summon_Top_Banner.png?version=a0ecb902dbf445c5de3615c370250d51"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 250, pickUpCards: []int{204, 206, 205, 156, 48}},
		{star: 5, cardType: 1, rareType: 1, Prob: 60},
		{star: 5, cardType: 2, rareType: 1, Prob: 90},
		{star: 4, Prob: 600, pickUpCards: []int{53, 59, 167}},
		{star: 4, cardType: 1, rareType: 1, Prob: 500},
		{star: 4, cardType: 2, rareType: 1, Prob: 500},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/b/bf/Monster_Hunter_Primal_Crisis_%28Part_One%29_Summon_Top_Banner.png?version=5d6a2bb75a78eb140a99c44578249111"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 250, pickUpCards: []int{207, 208, 210, 213, 214}},
		{star: 5, cardType: 1, rareType: 1, Prob: 60},
		{star: 5, cardType: 2, rareType: 1, Prob: 90},
		{star: 4, Prob: 800, pickUpCards: []int{209, 211, 212, 171}},
		{star: 4, cardType: 1, rareType: 1, Prob: 400},
		{star: 4, cardType: 2, rareType: 1, Prob: 400},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/6/6c/A_Clawful_Caper_%28Summon_Showcase%29_Summon_Top_Banner.png?version=a39e274e01dadbe1862de9175d917ce1"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 200, pickUpCards: []int{215, 216, 217, 154}},
		{star: 5, cardType: 1, rareType: 1, Prob: 80},
		{star: 5, cardType: 2, rareType: 1, Prob: 120},
		{star: 4, Prob: 800, pickUpCards: []int{218, 219, 220, 172}},
		{star: 4, cardType: 1, rareType: 1, Prob: 400},
		{star: 4, cardType: 2, rareType: 1, Prob: 400},
		{star: 3, cardType: 1, rareType: 1, Prob: 4800},
		{star: 3, cardType: 2, rareType: 1, Prob: 3200},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/e/e0/Dragonyule_Defenders_2_Summon_Top_Banner.png?version=d0101fc89559e39ab4b9f5cac218640a"})
	res = append(res, &cardMgrConfig{configs: []*CardSetConfig{
		{star: 5, Prob: 150, pickUpCards: []int{221, 222, 223}},
		{star: 5, cardType: 1, rareType: 1, Prob: 100},
		{star: 5, cardType: 2, rareType: 1, Prob: 150},
		{star: 4, Prob: 600, pickUpCards: []int{224, 225, 226}},
		{star: 4, cardType: 1, rareType: 1, Prob: 500},
		{star: 4, cardType: 2, rareType: 1, Prob: 500},
		{star: 3, Prob: 1000, pickUpCards: []int{227}},
		{star: 3, cardType: 1, rareType: 1, Prob: 4300},
		{star: 3, cardType: 2, rareType: 1, Prob: 2700},
	}, topBannerUrl: "https://gamepedia.cursecdn.com/dragalialost_gamepedia_en/6/64/Halloween_Fantasia_2_Summon_Top_Banner.png?version=dfd325b1480e28eda414fce369b09448"})

	return res
}
