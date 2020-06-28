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
	in           chan bool
	lastShoutOut map[int64]time.Time //记录每个群上一次发送消息的时间
}

type messageNode struct {
	f       func()
	groupID int64
}

func init() {

}

func (m *MessageQueue) Start() {
	m.list = list.List{}
	m.lastShoutOut = map[int64]time.Time{}
	m.in = make(chan bool)
	tickChan := time.Tick(time.Millisecond * 10)
	go func() {
		for {
			select {
			case _ = <-tickChan:
				m.sendOutOne()
			case _ = <-m.in:
				m.sendOutOne()
				//time.Sleep(time.Second * 1)
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
	node := m.list.Front().Value.(messageNode)

	lastTime := time.Now().Sub(m.lastShoutOut[node.groupID]).Seconds()
	//fmt.Println("sendOutOne", lastTime, time.Now())
	if lastTime < 1 {
		temp := m.list.Front()
		m.list.Remove(m.list.Front())
		m.list.PushBack(temp.Value)
		//fmt.Println("will not call")
		return
	}
	node.f()
	m.list.Remove(m.list.Front())
	m.lastShoutOut[node.groupID] = time.Now()
}

func (m *MessageQueue) addOne(f func(), groupid int64) {
	m.Lock()
	defer m.Unlock()
	m.list.PushBack(messageNode{
		f:       f,
		groupID: groupid,
	})
	select {
	case m.in <- true:
		fmt.Println("sent message")
	default:
		fmt.Println("queue is processing")
	}
}
