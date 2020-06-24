package cards

import "testing"

func TestCardCollectionV2_IsCardsExist(t *testing.T) {

	type fields struct {
		cardSets     []*CardSet
		TopBannerUrl string
	}
	type args struct {
		cards []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"testExist", fields{cardSets: []*CardSet{{
			cards: []Card{{
				ID: 1,
			}},
		}}}, args{cards: []int{1}}, true},
		{"testNotExist", fields{cardSets: []*CardSet{{
			cards: []Card{{
				ID: 1,
			}},
		}}}, args{cards: []int{2}}, false},
		{"testMultiExist", fields{cardSets: []*CardSet{{
			cards: []Card{{
				ID: 1,
			}, {
				ID: 2,
			}},
		}}}, args{cards: []int{1, 2}}, true},
		{"testMultiNotExist", fields{cardSets: []*CardSet{{
			cards: []Card{{
				ID: 1,
			}, {
				ID: 2,
			}},
		}}}, args{cards: []int{1, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CardCollectionV2{
				cardSets:     tt.fields.cardSets,
				TopBannerUrl: tt.fields.TopBannerUrl,
			}
			if got := c.IsCardsExist(tt.args.cards); got != tt.want {
				t.Errorf("IsCardsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
