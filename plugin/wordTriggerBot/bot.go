package wordTriggerBot

import (
	"iotqq-plugins-demo/Go/plugin"
	"math/rand"
	"sync"
	"time"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&wordTriggerBot{7, sync.Mutex{}})
}

type wordTriggerBot struct {
	priority int //[0~1000)
	mx       sync.Mutex
}

func (r *wordTriggerBot) Priority() int {
	return r.priority
}

func (r *wordTriggerBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if req.Content == "" && len(req.GroupPics) == 0 {
		return false, true
	}
	r.mx.Lock()
	defer r.mx.Unlock()
	for _, state := range stateList {
		//fmt.Println(state.response, state.coolDown, time.Since(state.lastTriggerTime))
		if state.triggerTimes > state.times || state.coolDown > time.Since(state.lastTriggerTime) {
			continue
		}
		if state.regex != nil && state.regex.MatchString(req.Content) {
			if rand.Intn(100) < state.probability {
				req.ExtraInfo = state.response
				state.triggerTimes++
				state.lastTriggerTime = time.Now()
				return true, false
			}
		}
		for _, pic := range req.GroupPics {
			if pic.FileMd5 == state.FileMd5 {
				req.ExtraInfo = state.response
				state.triggerTimes++
				state.lastTriggerTime = time.Now()
				return true, false
			}
		}
	}
	return false, true
}

func (r *wordTriggerBot) Process(req *plugin.Request) []*plugin.Result {
	res := []*plugin.Result{{Content: req.ExtraInfo.(string)}}
	return res
}
