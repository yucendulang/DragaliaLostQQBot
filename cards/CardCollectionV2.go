package cards

type CardCollectionV2 struct {
	cardSets     []*CardSet
	TopBannerUrl string
}

func (c CardCollectionV2) PickUpByStar(star int) []*CardSet {
	var res []*CardSet
	for _, set := range c.cardSets {
		if set.star == star {
			res = append(res, set)
		}
	}
	return res
}

func (c CardCollectionV2) IsCardsExist(cards []int) bool {
	m := make(map[int]bool)
	for _, card := range cards {
		m[card] = false
	}
	for i := range c.cardSets {
		for _, item := range c.cardSets[i].cards {
			if _, ok := m[item.ID]; ok {
				m[item.ID] = true
			}
		}
	}
	for _, card := range cards {
		if !m[card] {
			return false
		}
	}
	return true
}
