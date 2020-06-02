package summon

import (
	"fmt"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/userData"
	"math/rand"
	"testing"
	"time"
)

func TestTenSummon(t *testing.T) {
	tests := []struct {
		name    string
		wantRes SummonRecord
	}{
		{"basic", SummonRecord{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := GetMultiSummon(1000)(&userData.User{})
			totalWater := 0
			for _, card := range gotRes.Card {
				totalWater += card.Water
			}
			fmt.Printf(gotRes.Format())
			fmt.Println("Get water ", totalWater)
		})
	}
}

func TestTenSummonRate(t *testing.T) {
	tests := []struct {
		name    string
		wantRes SummonRecord
	}{
		{"basic", SummonRecord{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summonNum := 1000000
			rand.Seed(time.Now().Unix())
			start := time.Now()
			num1, num2, num3 := 0, 0, 0
			user := &userData.User{SummonCardNum: 10000000}
			for i := 0; i < summonNum; i++ {
				gotRes := TenSummon(user)
				for _, card := range gotRes.Card {
					if card.Star == 5 {
						num1++
					}
				}
			}
			period := time.Since(start)
			fmt.Println(num1, num2, num3, period.Milliseconds())
			prob := float64(num1) / float64(summonNum)
			if !(prob > 0.645 && prob < 0.65) {
				t.Errorf("TenSummon() probability = %v, want %v", prob, "0.645-0.65")
			}
		})
	}
}

//
//func TestTenSummonRate2(t *testing.T) {
//	tests := []struct {
//		name    string
//		wantRes SummonRecord
//	}{
//		{"basic", SummonRecord{}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rand.Seed(time.Now().Unix())
//			fmt.Println(main.SimAllIn(100, []int{191, 190, 189}))
//		})
//	}
//}
//
//func TestSingleSummon(t *testing.T) {
//	tests := []struct {
//		name    string
//		wantRes SummonRecord
//	}{
//		{"basic", SummonRecord{}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotRes := SingleSummon(&userData.User{}, false)
//			fmt.Println(gotRes.Format())
//			fmt.Println(len(cards.Cards))
//			for i, card := range cards.Cards {
//				if i+1 != card.ID {
//					fmt.Println("find wrong index", i)
//				}
//			}
//		})
//	}
//}

func TestSummonRecord_ImageFormat(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SummonRecord{
				Desc: "",
			}
			for i := 0; i < 10; i++ {
				res.Card = append(res.Card, SummonCard{
					Card: cards.Cards[1],
					New:  false,
				})
			}
			fmt.Println(res.ImageFormat(0, 0))
		})
	}
}
