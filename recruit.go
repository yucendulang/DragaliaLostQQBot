package main

import (
	"fmt"
	"iotqq-plugins-demo/Go/common"
	iotqq "iotqq-plugins-demo/Go/model"
	"iotqq-plugins-demo/Go/random"
	"strings"
	"time"
)

var selfTempQQ int64 = 2834323101

//循环招募的时间间隔
const period = 60

type Recruit struct {
	questID   int
	wantedNum int
	questName string
	qqgroupid int
	member    []Member
	close     chan bool
	shout     chan bool
}

type Member struct {
	QQ       int64
	Nickname string
}

type recruitManager []*Recruit

var manager recruitManager

func CreateRecruit(wantedNum int, questName string) *Recruit {
	i := 1
	for ; i < 100; i++ {
		var flag bool
		for _, recruit := range manager {
			if recruit.questID == i {
				flag = true
				break
			}
		}
		if !flag {
			break
		}
	}
	res := &Recruit{
		questID:   i,
		wantedNum: wantedNum + 1,
		questName: questName,
		close:     make(chan bool),
		shout:     make(chan bool),
	}
	manager = append(manager, res)
	return res
}

func GetRecruit(questid int) *Recruit {
	for i, i2 := range manager {
		if i2.questID == questid {
			return manager[i]
		}
	}
	return nil
}

func CancelAllRecruit() {
	for i, _ := range manager {
		manager[i].close <- false
	}
	manager = []*Recruit{}
}

func (r *Recruit) CancelRecruit(close bool) {
	for i, _ := range manager {
		if manager[i].questID == r.questID {
			manager[len(manager)-1], manager[i] = manager[i], manager[len(manager)-1]
			manager = manager[:len(manager)-1]
			r.close <- close
			return
		}
	}
	fmt.Println("It should not be here")
}

func (r *Recruit) ParticipateRecruit(member Member) {
	if r == nil {
		return
	}
	for _, m := range r.member {
		if m.QQ == member.QQ && m.QQ != 570966274 {
			iotqq.Send(r.qqgroupid, 2, m.Nickname+"不能重复参加"+random.RandomGetSuffix())
			return
		}
	}
	r.member = append(r.member, member)
	r.wantedNum--
	if r.wantedNum == 0 {
		fmt.Println("车队满人咯")
		r.CancelRecruit(true)
	} else if len(r.member) >= 2 {
		r.shout <- true
	}
}

func (r *Recruit) GetRecruitAd() string {
	res := fmt.Sprintf("%s招募中,缺%d人\n输入%d报名,输入c取消该车\n", r.questName, r.wantedNum, r.questID)
	var p []string
	for _, member := range r.member {
		p = append(p, member.Nickname)
	}
	res += fmt.Sprintf("现在参与的群友有:%s", strings.Join(p, ","))
	return res
}

func (r *Recruit) TryRecruit() {
	go func() {

		t := time.Tick(period * time.Second)
		res := r.GetRecruitAd()
		iotqq.Send(r.qqgroupid, 2, res)
	m:
		for {
			select {
			case _ = <-r.shout:
				r.shoutOut()
			case _ = <-t:
				r.shoutOut()
			case flag := <-r.close:
				fmt.Println("enter close logic")
				if flag {
					for _, member := range r.member {
						iotqq.Send2(r.qqgroupid, 2, r.questName+"发车"+random.RandomGetSuffix(), int(member.QQ))
						time.Sleep(time.Second * 2)
					}
				} else {
					iotqq.Send(r.qqgroupid, 2, "有内鬼!终止发车!")
				}
				break m
			}
		}
		fmt.Println("招募结束了呢")
	}()
}

func (r *Recruit) shoutOut() {
	res := r.GetRecruitAd()
	if common.HistoryRecord.IsExist(res, selfTempQQ) {
		fmt.Println("found recruit already")
		return
	}
	iotqq.Send(r.qqgroupid, 2, res)
}
