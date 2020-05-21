package cards

type Card struct {
	ID       int
	Star     int
	Title    string
	Water    int
	CardType int //1 is character 2 is dragon
	rareType int //1 is common summon
	IconUrl  string
}

func (c Card) PrintStar() string {
	var res string
	for i := 0; i < 5; i++ {
		if i < c.Star {
			res += "★"
		} else {
			res += "☆"
		}
	}
	return res
}
