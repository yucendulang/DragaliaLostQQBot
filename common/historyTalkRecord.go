package common

import (
	"container/list"
	"encoding/json"
	"fmt"
	"iotqq-plugins-demo/Go/util"
	"sync"
)

var HistoryRecord = historyTalkRecord{
	m: sync.Map{},
}

func init() {
	util.SaveRamVar("HistroyRecord", &HistoryRecord)
}

type historyTalkRecord struct {
	m sync.Map
}

type historyTalkRecordNode struct {
	threshold int
	List      *list.List
	mux       sync.RWMutex
}

func (h *historyTalkRecord) String() string {
	var nodes = make(map[int64][]node)
	h.m.Range(func(key, value interface{}) bool {
		nn := value.(*historyTalkRecordNode)
		gid := key.(int64)
		nn.mux.RLock()
		defer nn.mux.RUnlock()
		for e := nn.List.Front(); e != nil; e = e.Next() {
			nodes[gid] = append(nodes[gid], e.Value.(node))
		}
		return true
	})

	res, err := json.Marshal(nodes)
	//res,err:=json.Marshal(h)
	if err != nil {
		fmt.Println(err)
		return "{}"
	}
	return string(res)
}

type node struct {
	Content string
	Udid    int64
}

func (h *historyTalkRecord) Push(content string, udid int64, qqGroupID int64) {
	if value, ok := h.m.Load(qqGroupID); ok {
		nn := value.(*historyTalkRecordNode)
		nn.mux.Lock()
		defer nn.mux.Unlock()
		nn.List.PushBack(node{
			Content: content,
			Udid:    udid,
		})
		if nn.List.Len() > nn.threshold {
			nn.List.Remove(nn.List.Front())
		}
	} else {
		h.m.Store(qqGroupID, &historyTalkRecordNode{
			threshold: 3,
			List:      list.New(),
			mux:       sync.RWMutex{},
		})
	}

}

func (h *historyTalkRecord) IsExist(content string, udid, qqGroupID int64) bool {
	if value, ok := h.m.Load(qqGroupID); ok {
		nn := value.(*historyTalkRecordNode)
		nn.mux.RLock()
		defer nn.mux.RUnlock()
		for e := nn.List.Front(); e != nil; e = e.Next() {
			node := e.Value.(node)
			if node.Udid == udid && node.Content == content {
				return true
			}
		}
		return false
	} else {
		return false
	}
}
