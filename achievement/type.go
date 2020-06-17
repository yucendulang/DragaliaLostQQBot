package achievement

import "fmt"

type Achievement struct {
	ID      int
	Title   string
	Desc    string
	Trigger func(value interface{}) bool
}

func (a Achievement) Format(nickName string) string {
	return fmt.Sprintf("%s成就[%s]达成!", nickName, a.Title)
}

const (
	StickerFailed int = iota
	ReiceiveLotVolunter
	CoinMineRefresh
	GaChaThemAll
	SummonGreatThan10SSR
	SummonGreatThan20SSR
	SummonGreatThan30SSR
	SummonEqual0SSR
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
	{3, "一毛不拔也要全图鉴", "", nil},
	{4, "百连十虹抽卡我最行", "", func(value interface{}) bool {
		num := value.(int)
		if num >= 10 {
			return true
		} else {
			return false
		}
	}},
	{5, "欧皇再世百连二十虹", "", func(value interface{}) bool {
		num := value.(int)
		if num >= 20 {
			return true
		} else {
			return false
		}
	}},
	{6, "修炼几世百连三十虹", "", func(value interface{}) bool {
		num := value.(int)
		if num >= 30 {
			return true
		} else {
			return false
		}
	}},
	{7, "百连无虹大非才最欧", "", func(value interface{}) bool {
		num := value.(int)
		if num == 0 {
			return true
		} else {
			return false
		}
	}},
	//{3,"英文字母的顺序不影响阅读","",func(value interface{}) bool {
	//
	//},
}
