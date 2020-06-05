package achievement

import "fmt"

type Achievement struct {
	ID      int
	Title   string
	Desc    string
	Trigger func(value interface{}) bool
}

func (a Achievement) Format() string {
	return fmt.Sprintf("成就[%s]达成!", a.Title)
}

const (
	StickerFailed int = iota
	ReiceiveLotVolunter
	CoinMineRefresh
)

var AchievementList = []Achievement{
	{0, "成长轨迹即再次挑战", "", func(value interface{}) bool { return true }},
	{1, "此时此刻超越了太阳", "", func(value interface{}) bool {
		num := value.(int)
		if num > 7168 {
			return true
		} else {
			return false
		}
	}},
	{2, "水漫金币矿山的恐惧", "", nil},
	//{3,"英文字母的顺序不影响阅读","",func(value interface{}) bool {
	//
	//},
}
