package statisticsBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/plugin"
)

type PicNode struct {
	FileMD5 string
	Count   int
	Url     string
}

func init() {
	//statisticsBot := &statisticsBot{priority: 1, PicStatisticMap: make(map[string]*PicNode)}
	//plugin.FactoryInstance.RegisterPlugin(statisticsBot)
	//util.SaveRamVar("PicStatisticMap", statisticsBot)
}

type statisticsBot struct {
	priority        int //[0~1000)
	PicStatisticMap map[string]*PicNode
}

func (s *statisticsBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if len(req.GroupPics) != 0 {
		return true, true
	}
	return false, true
}

func (s *statisticsBot) Process(req *plugin.Request) []*plugin.Result {
	for _, pic := range req.GroupPics {
		if value, ok := s.PicStatisticMap[pic.FileMd5]; ok {
			value.Count++
		} else {
			s.PicStatisticMap[pic.FileMd5] = &PicNode{
				FileMD5: pic.FileMd5,
				Count:   1,
				Url:     pic.Url,
			}
		}
	}
	return nil
}

func (s *statisticsBot) Priority() int {
	return s.priority
}

func (s *statisticsBot) String() string {
	res := ""
	for key, value := range s.PicStatisticMap {
		res += fmt.Sprintf("key:%s,Count:%d,Url:%s", key, value.Count, value.Url)
	}
	return res
}
