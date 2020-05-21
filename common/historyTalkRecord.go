package common

import (
	"container/list"
	"encoding/json"
	"expvar"
	"fmt"
	"sync"
)

var HistoryRecord = historyTalkRecord{
	threshold: 3,
	List:      list.New(),
}

func init() {
	expvar.Publish("HistroyRecord", &HistoryRecord)
}

type historyTalkRecord struct {
	threshold  int
	List       *list.List
	lastRepeat string
	mux        sync.RWMutex
}

func (h *historyTalkRecord) String() string {
	h.mux.RLock()
	defer h.mux.RUnlock()
	var nodes []node
	for e := h.List.Front(); e != nil; e = e.Next() {
		nodes = append(nodes, e.Value.(node))
	}
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

func (h *historyTalkRecord) Push(content string, udid int64) {
	h.mux.Lock()
	defer h.mux.Unlock()
	h.List.PushBack(node{
		Content: content,
		Udid:    udid,
	})
	if h.List.Len() > 3 {
		h.List.Remove(h.List.Front())
	}
}

func (h *historyTalkRecord) IsExist(content string, udid int64) bool {
	h.mux.RLock()
	defer h.mux.RUnlock()
	for e := h.List.Front(); e != nil; e = e.Next() {
		node := e.Value.(node)
		if node.Udid == udid && node.Content == content {
			return true
		}
	}
	return false
}
