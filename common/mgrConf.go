package common

import (
	"encoding/json"
	"iotqq-plugins-demo/Go/util"
	"sync"
)

func init() {
	util.SaveRamVar("GroupMgrConf", &GroupMgrConf)
}

var GroupMgrConf = GroupMgrConfType{
	GroupMap: map[int64]*GroupMgrConfItem{},
}

type GroupMgrConfType struct {
	GroupMap map[int64]*GroupMgrConfItem
	mux      sync.RWMutex
}

type GroupMgrConfItem struct {
	IsIntroOn   bool
	IsBotOn     bool
	ManagerUDID int64
}

func (g *GroupMgrConfType) String() string {
	res := ""
	g.mux.RLock()
	defer g.mux.RUnlock()
	for _, item := range g.GroupMap {
		js, _ := json.Marshal(item)
		res += string(js)
	}
	return res
}

func (g *GroupMgrConfType) GetByGroupID(id int64) *GroupMgrConfItem {
	return g.read(id)
}

func (g *GroupMgrConfType) IsBotOn(id int64) bool {
	return g.read(id).IsBotOn
}

func (g *GroupMgrConfType) read(id int64) *GroupMgrConfItem {
	if res, ok := g.GroupMap[id]; ok {
		return res
	} else {
		return &GroupMgrConfItem{
			IsIntroOn: false,
			IsBotOn:   false,
		}
	}
}

func (g *GroupMgrConfType) SetIsBotOn(id int64, b bool) {
	if res, ok := g.GroupMap[id]; ok {
		res.IsBotOn = b
	}
}

func (g *GroupMgrConfType) IsManager(id int64, id2 int64) bool {
	if g.read(id).ManagerUDID == id2 {
		return true
	} else {
		return false
	}
}
