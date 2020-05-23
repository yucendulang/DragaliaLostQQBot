package model

import (
	"container/list"
	"sync"
	"time"
)

type MessageQueue struct {
	list list.List
	sync.Mutex
}

func init() {

}

func (m *MessageQueue) Start() {
	m.list = list.List{}
	go Periodlycall(time.Second, m.sendOutOne)
}

func (m *MessageQueue) sendOutOne() {
	m.Lock()
	defer m.Unlock()
	if m.list.Len() == 0 {
		return
	}
	f := m.list.Front().Value.(func())
	f()
	m.list.Remove(m.list.Front())
}

func (m *MessageQueue) addOne(f func()) {
	m.Lock()
	defer m.Unlock()
	m.list.PushBack(f)
}
