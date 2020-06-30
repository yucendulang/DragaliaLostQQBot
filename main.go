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
	_ "iotqq-plugins-demo/Go/plugin/summonGame/rebornBot"
	_ "iotqq-plugins-demo/Go/plugin/summonGame/staticQueryBot"
	_ "iotqq-plugins-demo/Go/plugin/wordTriggerBot"
	"iotqq-plugins-demo/Go/random"
	"iotqq-plugins-demo/Go/summon"
	"iotqq-plugins-demo/Go/userData"
	"iotqq-plugins-demo/Go/util"
	"log"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var url1 string
var site = "127.0.0.1"
var port = 8888

func main() {
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			fmt.Println("参数", i, "是", arg)
		}
		common.QQ = os.Args[1]
		i, _ := strconv.Atoi(os.Args[1])
		common.QQInt = int64(i)
	}

	rand.Seed(time.Now().Unix())
	model.StartPicServer(true)

	//return

	userData.MaxCollectionNum = len(cards.Cards)
	common.FiveStarCharacterNum = cards.GetCardsNumByStarType(5, 1)
	common.FourStarCharacterNum = cards.GetCardsNumByStarType(4, 1)
	common.ThreeStarCharacterNum = cards.GetCardsNumByStarType(3, 1)
	common.FiveStarDragonNum = cards.GetCardsNumByStarType(5, 2)
	common.FourStarDragonNum = cards.GetCardsNumByStarType(4, 2)
	common.ThreeStarDragonNum = cards.GetCardsNumByStarType(3, 2)
	common.ThreeStarDragonNum = cards.GetCardsNumByStarType(3, 2)
	common.GachaPoolNum = cards.GetCardsNumByrareType([]int{cards.RareTypePermanent, cards.RareTypeNotPermanent, cards.RareTypeGala})
	common.NoGachaPoolNum = cards.GetCardsNumByrareType([]int{cards.RareTypeStory, cards.RareTypeEvent})

	userData.UserDataLoad()
	util.SignalNotify(userData.UserDataSave)
	util.RestoreRamVar()
	summon.InitImageSource()
	model.StartTick()
	mq := model.MessageQueue{}
	mq.Start()
	recruitexp := regexp.MustCompile("(.*)招募(.*)缺([0-9])")
	recruitCanjiaExp := regexp.MustCompile("^[0-9]$")
	buildCommand := regexp.MustCompile("\"(?:@修玛吉亚-Du|@矛盾的人偶) 建造(.*?)\"")

	url1 = site + ":" + strconv.Itoa(port)
	model.Set(url1, common.QQ, &mq)
	runtime.GOMAXPROCS(runtime.NumCPU())

	plugin.FactoryInstance.SetOptions(common.QQInt)

	go func() {

		for {
			connect(buildCommand, recruitexp, recruitCanjiaExp)
			time.Sleep(time.Second * 5)
		}
	}()

	//model.Periodlycall(60*time.Second, userData.UserDataSave)
	model.Periodlycall(60*time.Second, func() {
		fmt.Println("Server 60s tick ", time.Now().String())
	})
	//fmt.Println(" [x] Complete")
}

func connect(buildCommand *regexp.Regexp, recruitexp *regexp.Regexp, recruitCanjiaExp *regexp.Regexp) {
	var fail = make(chan bool)
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(site, port, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On("OnGroupMsgs", func(h *gosocketio.Channel, args model.Message) {
		//if args.CurrentPacket.Data.FromUserID != 570966274 {
		//	return
		//}
		processGroupMsg(args, buildCommand, recruitexp, recruitCanjiaExp)

	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On("OnFriendMsgs", func(h *gosocketio.Channel, args model.Message) {
		var mess model.Data = args.CurrentPacket.Data
		log.Printf("私聊消息:%+v", mess)
		if mess.FromUin != 570966274 {
			return
		}
		mess.FromUserID = mess.FromUin
		plugin.FactoryInstance.Run(mess, func(content, picUrl string) {
			if picUrl != "" {
				model.SendPic(int(mess.FromUin), 1, content, picUrl)
			} else {
				model.Send(int(mess.FromUin), 1, content)
			}
		})

	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		fmt.Println("Disconnected")
		fail <- true
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		fmt.Println("连接成功")
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1 * time.Second)
	if !model.SendJoin(c, common.QQ) {
		fmt.Println("login failed.Retry...")
		return
	}

	_ = <-fail
	fmt.Println("some thing happen failed exit connect")
}

func processGroupMsg(args model.Message, buildCommand *regexp.Regexp, recruitexp *regexp.Regexp, recruitCanjiaExp *regexp.Regexp) {

	var mess model.Data = args.CurrentPacket.Data

	if !common.GroupMgrConf.IsBotOn(int64(mess.FromGroupID)) {
		if strings.Contains(mess.Content, ".bot on") && common.GroupMgrConf.IsManager(int64(mess.FromGroupID), mess.FromUserID) {
			common.GroupMgrConf.GetByGroupID(int64(mess.FromGroupID)).IsBotOn = true
			model.Send(mess.FromGroupID, 2, "bot on")
		}
		return
	} else {
		if strings.Contains(mess.Content, ".bot off") && common.GroupMgrConf.IsManager(int64(mess.FromGroupID), mess.FromUserID) {
			common.GroupMgrConf.GetByGroupID(int64(mess.FromGroupID)).IsBotOn = false
			model.Send(mess.FromGroupID, 2, "bot off")
		}
	}

	common.HistoryRecord.Push(mess.Content, mess.FromUserID, int64(mess.FromGroupID))

	if mess.FromUserID == common.QQInt {
		return
	}
	/*
		mess.Content 消息内容 string
		mess.FromGroupID 来源QQ群 int
		mess.FromUserID 来源QQ int64
		mess.iotqqType 消息类型 string
	*/
	nickName := util.FixName(mess.FromNickName)
	fmt.Println("群聊消息: ", mess.FromGroupID, nickName+"<"+strconv.FormatInt(mess.FromUserID, 10)+">: "+mess.Content)

	if mess.FromUserID == 570966274 && util.KeyWordTrigger(mess.Content, "repeat") {
		str := ""
		for i := 0; i < 5005; i++ {
			str += "囧"
			model.Send(mess.FromGroupID, 2, fmt.Sprintf("%d%s", i, str))
		}
	}

	//if util.KeyWordTrigger(mess.Content, "abcd all") {
	//	userData.UserRange(func(key, value interface{}) bool {
	//		value.(*userData.User).SummonCardNum += 200
	//		return true
	//	})
	//}
	//
	//if util.KeyWordTrigger(mess.Content, "abcd coinmine") {
	//	userData.UserRange(func(key, value interface{}) bool {
	//		//value.(*userData.User).BuildIndex = append(value.(*userData.User).BuildIndex, common.BuildRecord{Index: 2, Level: 1})
	//		index := -1
	//		for i := range value.(*userData.User).AchievementList {
	//			if value.(*userData.User).AchievementList[i].Index == achievement.SummonGreatThan20SSR {
	//				index = i
	//			}
	//		}
	//		if index == -1 {
	//			return true
	//		}
	//		value.(*userData.User).AchievementList = append(value.(*userData.User).AchievementList[:index], value.(*userData.User).AchievementList[index+1:]...)
	//		return true
	//	})
	//}

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
			userData.SaveUserByUDID(mess.FromUserID)
			model.Send(mess.FromGroupID, 2, nickName+out+"花费"+strconv.Itoa(cost)+"💧")
		} else {
			model.Send(mess.FromGroupID, 2, nickName+out)
		}
		return
	}

	rec := recruitexp.FindStringSubmatch(mess.Content)
	if len(rec) > 0 {
		fmt.Println("start recruit")
		num, _ := strconv.Atoi(rec[3])
		questName := rec[2]
		if len(questName) == 0 {
			questName = rec[1]
		}
		recruit := CreateRecruit(num, questName)
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
		if r != nil && r.qqgroupid == mess.FromGroupID {
			r.ParticipateRecruit(Member{
				QQ:       mess.FromUserID,
				Nickname: nickName,
			})
		}
	}

	if mess.Content == "c" {
		CancelAllRecruit(mess.FromUserID)
	}

	//if mess.Content == "testrapid" {
	//	for i := 0; i < 10; i++ {
	//		model.Send(992028272, 2, fmt.Sprint("echo back", i))
	//		model.Send(1128869023, 2, fmt.Sprint("echo back", i))
	//	}
	//
	//}

	plugin.FactoryInstance.Run(mess, func(content, picUrl string) {
		if picUrl != "" {
			model.SendPic(mess.FromGroupID, 2, content, picUrl)
		} else {
			model.Send(mess.FromGroupID, 2, content)
		}
	})
}
