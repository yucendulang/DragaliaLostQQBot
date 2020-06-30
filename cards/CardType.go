package cards

type Card struct {
	ID       int
	Star     int
	Title    string
	Water    int
	CardType int //1 is character 2 is dragon
	RareType int //1-common(PERMANENT 2-限定(NOTPERMANENT 3-fes(GALA 4-story 5-event
	IconUrl  string
}

type RareType int

const (
	RareTypePermanent int = iota + 1
	RareTypeNotPermanent
	RareTypeGala
	RareTypeStory
	RareTypeEvent
	RareTypeCount
)

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
