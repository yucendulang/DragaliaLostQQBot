package probabilityCalBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"strings"
	"time"
)

func SimAllIn(num int, card []int) string {
	start := time.Now()
	num1 := 0
	max := 10000000 / num
	summonNum := num / 10
	for i := 0; i < max; i++ {
		user := &userData.User{}
		for j := 0; j < summonNum; j++ {
			summon.TenSummon(user)
		}
		gachaThemAll := true
		for _, cardIndex := range card {
			if !util.IntContain(cardIndex, user.CardIndex) {
				gachaThemAll = false
			}
		}
		if gachaThemAll {
			num1++
		}
	}
	period := time.Since(start)
	fmt.Printf("本次模拟耗时%ds\n", period.Milliseconds()/1000)
	return fmt.Sprintf("进行%d次模拟,简化模型进行%d次十连召唤\n满足目标的概率为%.2f%%",
		max, summonNum, float32(num1)/float32(max)*100)

}

func SimMustGet(card []int) string {
	return "not implement"
}

func SimMustGetV2(card []int) string {
	start := time.Now()
	max := 100000
	sum := 0
	for i := 0; i < max; i++ {
		user := &userData.User{}
		gachaThemAll := false
		num := 10
		for ; !gachaThemAll; num += 10 {
			//todo
			summon.TenSummonByCollection(user, cards.CardMgr.PickUp(1))
			gachaThemAll = true
			for _, cardIndex := range card {
				if !util.IntContain(cardIndex, user.CardIndex) {
					gachaThemAll = false
				}
			}
		}
		sum += num
	}
	period := time.Since(start)
	fmt.Printf("本次模拟耗时%ds\n", period.Milliseconds()/1000)
	return fmt.Sprintf("进行%d次模拟,average sum count is %d",
		max, int(float32(sum)/float32(max)))

}

func SimParse(cardNames string, drawNum int) (string, func() string, error) {
	names := strings.Split(cardNames, ",")
	cards := cards.FindCardIndex(names)
	for i := range cards {
		if cards[i] == -1 {
			return "", nil, fmt.Errorf("找不到%s,敬爱的殿下请检查是否输入错召唤对象%s", names[i], random.RandomGetSuffix())
		}
	}

	num := drawNum
	if num%10 != 0 {
		return "", nil, fmt.Errorf("暂时只支持十连,%d不是10的倍数", num)
	}

	var f func() string
	if num == 0 {
		f = func() string {
			out := fmt.Sprintf("敬爱的殿下,我抽晕了%s\n", random.RandomGetSuffix())
			out += SimMustGet(cards)
			return out
		}
	} else {
		f = func() string {
			allIn := SimAllIn(num, cards)
			out := fmt.Sprintf("敬爱的殿下,我抽晕了%s\n", random.RandomGetSuffix())
			out += allIn
			return out
		}
	}

	return fmt.Sprintf("敬爱的殿下,请耐心等待结果,正在努力计算%s...", random.RandomGetSuffix()), f, nil
}

func SimSSRGet(num int) string {
	start := time.Now()
	num15, num14, num13 := 0, 0, 0
	max := 10000000 / num
	summonNum := num / 10
	user := &userData.User{}
	for i := 0; i < max; i++ {
		summonRecord := summon.GetMultiSummon(100)(user)
		for _, index := range summonRecord.Card {
			if index.Star == 5 {
				num15++
			} else if index.Star == 4 {
				num14++
			} else if index.Star == 3 {
				num13++
			}
		}
	}
	period := time.Since(start)
	fmt.Printf("本次模拟耗时%ds\n", period.Milliseconds()/1000)
	return fmt.Sprintf("进行%d次模拟,简化模型进行%d次十连召唤\n,%dSSR,%dSR,%dR,%dTotal",
		max, summonNum, num15, num14, num13, num15+num14+num13)
}

func SimMustGet2() string {
	card := []int{189, 190, 191}
	start := time.Now()
	max := 100000
	sum := 0
	resultMap := map[int]int{}
	for i := 0; i < max; i++ {
		user := &userData.User{}
		card1Get, card2Get, card3Get := 0, 0, 0
		gachaThemAll := false
		num := 10
		for ; !gachaThemAll; num += 10 {
			summonRecord := summon.TenSummon(user)
			for _, summonCard := range summonRecord.Card {
				switch summonCard.ID {
				case card[0]:
					card1Get++
				case card[1]:
					card2Get++
				case card[2]:
					card3Get++
				}
			}

			gachaThemAll = true
			for _, cardIndex := range card {
				if !util.IntContain(cardIndex, user.CardIndex) {
					gachaThemAll = false
				}
			}
		}
		resultMap[card1Get*10000+card2Get*100+card3Get] += 1
		sum += num
	}
	period := time.Since(start)
	fmt.Printf("本次模拟耗时%ds\n", period.Milliseconds()/1000)
	out := fmt.Sprintf("进行%d次模拟,average sum count is %d\n",
		max, int(float32(sum)/float32(max)))
	for key, value := range resultMap {
		out += fmt.Sprintf("%d,%d,%d,%d\n", key/10000, key%10000/100, key%100, value)
	}
	return out
}

func SimAllIn2(num int, card []int) string {
	start := time.Now()
	num1 := 0
	max := 100000
	summonNum := num / 10
	resultMap := map[int]int{}

	for i := 0; i < max; i++ {
		user := &userData.User{}
		cardGetNum := make([]int, len(card))
		for j := 0; j < summonNum; j++ {
			rc := summon.TenSummon(user)
			for _, summonCard := range rc.Card {
				for ic, c := range card {
					if c == summonCard.ID {
						cardGetNum[ic]++
					}
				}
			}
		}
		gachaThemAll := true
		for _, cardIndex := range card {
			if !util.IntContain(cardIndex, user.CardIndex) {
				gachaThemAll = false
			}
		}
		if gachaThemAll {
			num1++
		}
		resultMap[cardGetNum[0]*10000+cardGetNum[1]*100+cardGetNum[2]]++
	}
	period := time.Since(start)
	fmt.Printf("本次模拟耗时%ds\n", period.Milliseconds()/1000)
	out := fmt.Sprintf("进行%d次模拟,简化模型进行%d次十连召唤\n满足目标的概率为%.2f%s\n",
		max, summonNum, float32(num1)/float32(max)*100, "%")

	for key, value := range resultMap {
		out += fmt.Sprintf("%d,%d,%d,%d\n", key/10000, key%10000/100, key%100, value)
	}
	return out

}
