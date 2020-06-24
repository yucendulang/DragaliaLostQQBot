package probabilityCalBot

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSimAllIn(t *testing.T) {
	rand.Seed(time.Now().Unix())
	type args struct {
		num  int
		card []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"basic", args{
			num:  300,
			card: []int{189},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimAllIn(tt.args.num, 0, tt.args.card)
			fmt.Println(got)
		})
	}
}

func TestSimSSRGet(t *testing.T) {
	rand.Seed(time.Now().Unix())
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
	}{
		{"basic", args{num: 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimAllIn2(290, []int{189, 190, 191})
			fmt.Printf(got)
		})
	}
}

func TestSimMustGet(t *testing.T) {
	type args struct {
		card          []int
		cardPoolIndex int
	}
	tests := []struct {
		name string
		args args
	}{
		{"basic", args{card: []int{20}, cardPoolIndex: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimMustGet(tt.args.card, tt.args.cardPoolIndex)
			fmt.Println(got)
		})
	}
}
