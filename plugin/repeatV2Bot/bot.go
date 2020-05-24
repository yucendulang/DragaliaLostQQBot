package repeatV2Bot

import (
	"container/list"
	"iotqq-plugins-demo/Go/plugin"
	"iotqq-plugins-demo/Go/util"
	"sync"
)

func init() {
	v2 := &repeatV2Bot{priority: 997, threshold: 3}
	plugin.FactoryInstance.RegisterPlugin(v2)
	v2.list = list.New()
}

type repeatV2Bot struct {
	threshold  int //复读的阈值
	priority   int //[0~1000)
	list       *list.List
	lastRepeat string
	mux        sync.Mutex
}

type node struct {
	content string
	udid    int64
}

func (r *repeatV2Bot) Priority() int {
	return r.priority
}

func (r *repeatV2Bot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if req.Content == "" {
		return false, true
	}
	r.mux.Lock()
	defer r.mux.Unlock()
	r.list.PushBack(node{req.Content, req.Udid})
	if r.list.Len() > r.threshold {
		r.list.Remove(r.list.Front())
	} else {
		return false, true
	}

	if isRepeat(r) {
		r.lastRepeat = r.list.Front().Value.(node).content
		r.list.Init()
		return true, true
	} else {
		return false, true
	}
}

func isRepeat(r *repeatV2Bot) bool {
	var user []int64
	for e := r.list.Front(); e != nil; e = e.Next() {
		n := e.Value.(node)
		// 如果是单人复读
		if util.Int64Contain(n.udid, user) && n.udid != 570966274 {
			return false
		}
		user = append(user, n.udid)

		if e.Prev() == nil {
			continue
		} else {
			//没有复读
			if e.Prev().Value.(node).content != n.content {
				return false
			}
		}
	}
	return true
}

func (r *repeatV2Bot) Process(_ *plugin.Request) []*plugin.Result {
	res := []*plugin.Result{{Content: r.lastRepeat}}
	return res
}
