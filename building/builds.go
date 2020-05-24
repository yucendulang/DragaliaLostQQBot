package building

import (
	"fmt"
	"iotqq-plugins-demo/Go/common"
)

type building struct {
	Cost  int
	index int
	f     func(*buildingEffect, int)
	Title string
}

var BuildList []building

func init() {
	new(100000, 0, "觉醒之岚树", func(b *buildingEffect, level int) {
		b.RepeatBonus += 10 * level
	})
	new(200000, 1, "水祭坛", func(b *buildingEffect, level int) {
		b.RepeatProbability += 10 * level
	})
	new(200000, 2, "金币矿山", func(b *buildingEffect, level int) {
		b.VolunterMineProduct = 50 * (100 + (level-1)*10) / 100
	})
	//new(100000, 3, "水祭坛2", func(b *buildingEffect, level int) {
	//	b.RepeatProbability += 10 * level
	//})
}

func new(c, index int, title string, f func(*buildingEffect, int)) {
	BuildList = append(BuildList, building{
		Cost:  c,
		index: index,
		f:     f,
		Title: title,
	})
}

func ConstructNewBuilding(name string) (string, int) {
	index, ok := FindBuildingByName(name)
	if !ok {
		return fmt.Sprintf("没有找到名字为%s的建筑", name), -1
	}
	return fmt.Sprintf("%s建造完成", name), index
}

func FindBuildingByName(name string) (int, bool) {
	for i := range BuildList {
		if BuildList[i].Title == name {
			return BuildList[i].index, true
		}
	}
	return -1, false
}

func GetBuildEffect(record []common.BuildRecord) *buildingEffect {
	eff := &buildingEffect{
		RepeatProbability:   0,
		RepeatBonus:         100,
		VolunterMineProduct: 0,
	}
	for _, r := range record {
		BuildList[r.Index].f(eff, r.Level)
	}
	return eff
}
