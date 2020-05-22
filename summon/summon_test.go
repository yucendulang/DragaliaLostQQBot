package summon

import (
	"fmt"
	"iotqq-plugins-demo/Go/cards"
	"testing"
)

//
//func TestTenSummon(t *testing.T) {
//	tests := []struct {
//		name    string
//		wantRes SummonRecord
//	}{
//		{"basic", SummonRecord{}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotRes := TenSummon(&userData.User{})
//			fmt.Printf(gotRes.Format())
//		})
//	}
//}
//
//func TestTenSummonRate(t *testing.T) {
//	tests := []struct {
//		name    string
//		wantRes SummonRecord
//	}{
//		{"basic", SummonRecord{}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rand.Seed(time.Now().Unix())
//			start := time.Now()
//			num1, num2, num3 := 0, 0, 0
//			user := &userData.User{SummonCardNum: 10000000}
//			for i := 0; i < 10000000; i++ {
//				gotRes := TenSummon(user)
//				for _, card := range gotRes.Card {
//					if card.ID == 191 {
//						num1++
//					} else if card.ID == 190 {
//						num2++
//					} else if card.ID == 189 {
//						num3++
//					}
//				}
//			}
//			period := time.Since(start)
//			fmt.Println(num1, num2, num3, period.Milliseconds())
//		})
//	}
//}
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
