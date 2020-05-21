package building

type buildingEffect struct {
	RepeatProbability int
	RepeatBonus       int
}

func (b *buildingEffect) GetExtraRepeatBonus() float32 {
	return float32(b.RepeatBonus) / 100
}
