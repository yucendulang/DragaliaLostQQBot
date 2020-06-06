package main

import (
	"fmt"
	"iotqq-plugins-demo/Go/building"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/common"
	"iotqq-plugins-demo/Go/model"
	"iotqq-plugins-demo/Go/plugin"
	_ "iotqq-plugins-demo/Go/plugin/helpBot"
	_ "iotqq-plugins-demo/Go/plugin/repeatBot"
	_ "iotqq-plugins-demo/Go/plugin/repeatV2Bot"
	_ "iotqq-plugins-demo/Go/plugin/statisticsBot"
	_ "iotqq-plugins-demo/Go/plugin/stickerBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/announceBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/collectorBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/gachaBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/probabilityCalBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/queryBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/staticQueryBot"
	_ "iotqq-plugins-demo/Go/plugin/wordTriggerBot"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"log"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var url1, qq string
var site = "127.0.0.1"
var port = 8888

func main() {
	model.StartPicServer(true)

	//return

	userData.MaxCollectionNum = len(cards.Cards)
	common.FiveStarCharacterNum = cards.GetCardsNumByStarType(5, 1)
	common.FourStarCharacterNum = cards.GetCardsNumByStarType(4, 1)
	common.ThreeStarCharacterNum = cards.GetCardsNumByStarType(3, 1)
	common.FiveStarDragonNum = cards.GetCardsNumByStarType(5, 2)
	common.FourStarDragonNum = cards.GetCardsNumByStarType(4, 2)
	common.ThreeStarDragonNum = cards.GetCardsNumByStarType(3, 2)

	userData.UserDataLoad()
	util.SignalNotify()
	util.RestoreRamVar()
	summon.InitImageSource()
	model.StartTick()
	mq := model.MessageQueue{}
	mq.Start()
	recruitexp := regexp.MustCompile("招募(.*)缺([0-9])")
	recruitCanjiaExp := regexp.MustCompile("^[0-9]$")
	buildCommand := regexp.MustCompile("\"@修玛吉亚-Du 建造(.*?)\"")

	qq = "2834323101"
	url1 = site + ":" + strconv.Itoa(port)
	model.Set(url1, qq, &mq)
	runtime.GOMAXPROCS(runtime.NumCPU())

	qqInt, _ := strconv.Atoi(qq)
	plugin.FactoryInstance.SetOptions(int64(qqInt))

	go func() {
		for {
			connect(buildCommand, recruitexp, recruitCanjiaExp)
			time.Sleep(time.Second * 5)
		}
	}()

	model.Periodlycall(60*time.Second, userData.UserDataSave)
	//log.Println(" [x] Complete")
}

func connect(buildCommand *regexp.Regexp, recruitexp *regexp.Regexp, recruitCanjiaExp *regexp.Regexp) {
	var fail = make(chan bool)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(site, port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Println(err)
		return
	}
	err = c.On("OnGroupMsgs", func(h *gosocketio.Channel, args model.Message) {
		processGroupMsg(args, buildCommand, recruitexp, recruitCanjiaExp)

	})
	if err != nil {
		log.Println(err)
		return
	}
	err = c.On("OnFriendiotqqs", func(h *gosocketio.Channel, args model.Message) {
		var mess model.Data = args.CurrentPacket.Data
		log.Printf("私聊消息:%+v", mess)
		if mess.FromUserID != 570966274 {
			return
		}
		plugin.FactoryInstance.Run(mess)

	})
	if err != nil {
		log.Println(err)
		return
	}
	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Println("Disconnected")
		fail <- true
	})
	if err != nil {
		log.Println(err)
		return
	}
	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("连接成功")
	})
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(1 * time.Second)
	if !model.SendJoin(c, qq) {
		fmt.Println("login failed.Retry...")
		return
	}

	_ = <-fail
	fmt.Println("some thing happen failed exit connect")
}

func processGroupMsg(args model.Message, buildCommand *regexp.Regexp, recruitexp *regexp.Regexp, recruitCanjiaExp *regexp.Regexp) {
	rand.Seed(time.Now().Unix())
	var mess model.Data = args.CurrentPacket.Data

	common.HistoryRecord.Push(mess.Content, mess.FromUserID)

	if q, _ := strconv.Atoi(qq); mess.FromUserID == int64(q) {
		return
	}
	/*
		mess.Content 消息内容 string
		mess.FromGroupID 来源QQ群 int
		mess.FromUserID 来源QQ int64
		mess.iotqqType 消息类型 string
	*/
	nickName := util.FixName(mess.FromNickName)
	log.Println("群聊消息: ", mess.FromGroupID, nickName+"<"+strconv.FormatInt(mess.FromUserID, 10)+">: "+mess.Content)

	if util.KeyWordTrigger(mess.Content, "abcd all") {
		userData.UserRange(func(key, value interface{}) bool {
			value.(*userData.User).SummonCardNum += 200
			return true
		})
	}

	if util.KeyWordTrigger(mess.Content, "abcd coinmine") {
		userData.UserRange(func(key, value interface{}) bool {
			value.(*userData.User).BuildIndex = append(value.(*userData.User).BuildIndex, common.BuildRecord{Index: 2, Level: 1})
			return true
		})
	}

	buildComm := buildCommand.FindStringSubmatch(mess.Content)
	if len(buildComm) > 0 {
		out, index := building.ConstructNewBuilding(buildComm[1])
		if index >= 0 {
			user := userData.GetUser(mess.FromUserID)
			var level int
			var levelIndex int
			for i, buildIndex := range user.BuildIndex {
				if buildIndex.Index == index {
					level = buildIndex.Level
					levelIndex = i
					break
				}
			}
			cost := building.BuildList[index].Cost * level
			if user.Water < cost {
				model.Send(mess.FromGroupID, 2, fmt.Sprintf(nickName+"建造费用%d💧不够"+random.RandomGetSuffix(), cost))
				return
			} else {
				if level == 0 {
					user.BuildIndex = append(user.BuildIndex, common.BuildRecord{Index: index, Level: 1})
				} else {
					user.BuildIndex[levelIndex].Level++
				}

				user.Water -= cost
			}
			model.Send(mess.FromGroupID, 2, nickName+out+"花费"+strconv.Itoa(cost)+"💧")
		} else {
			model.Send(mess.FromGroupID, 2, nickName+out)
		}
		return
	}

	rec := recruitexp.FindStringSubmatch(mess.Content)
	if len(rec) > 0 {
		fmt.Println("start recruit")
		num, _ := strconv.Atoi(rec[2])
		recruit := CreateRecruit(num, rec[1])
		recruit.qqgroupid = mess.FromGroupID
		recruit.ParticipateRecruit(Member{
			QQ:       mess.FromUserID,
			Nickname: nickName,
		})
		recruit.TryRecruit()
		for _, s := range rec {
			fmt.Println(s)
		}
	}

	if recruitCanjiaExp.MatchString(mess.Content) {
		fmt.Println("有人参加任务")
		i, _ := strconv.Atoi(mess.Content)
		r := GetRecruit(i)
		r.ParticipateRecruit(Member{
			QQ:       mess.FromUserID,
			Nickname: nickName,
		})
	}

	if mess.Content == "c" {
		CancelAllRecruit(mess.FromUserID)
	}

	if mess.Content == "testrapid" {
		model.Send(mess.FromGroupID, 2, "echo back")
		model.Send(mess.FromGroupID, 2, "echo back")
		model.Send(mess.FromGroupID, 2, "echo back")
	}

	plugin.FactoryInstance.Run(mess)
}
