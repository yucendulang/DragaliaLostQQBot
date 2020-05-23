package model

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type MessageQueue struct {
	list list.List
	sync.Mutex
	in chan bool
}

func init() {

}

func (m *MessageQueue) Start() {
	m.list = list.List{}
	m.in = make(chan bool)
	tickChan := time.Tick(time.Second)
	go func() {
		for {
			select {
			case _ = <-tickChan:
				m.sendOutOne()
			case _ = <-m.in:
				m.sendOutOne()
				time.Sleep(time.Second * 1)
			}
		}
	}()
}

func (m *MessageQueue) sendOutOne() {
	m.Lock()
	defer m.Unlock()
	if m.list.Len() == 0 {
		return
	}
	f := m.list.Front().Value.(func())
	fmt.Println("sendOutOne", time.Now())
	f()
	m.list.Remove(m.list.Front())
}

func (m *MessageQueue) addOne(f func()) {
	m.Lock()
	defer m.Unlock()
	m.list.PushBack(f)
	select {
	case m.in <- true:
		fmt.Println("sent message")
	default:
		fmt.Println("queue is processing")
	}
}
