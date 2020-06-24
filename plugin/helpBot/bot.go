package helpBot

import (
	"iotqq-plugins-demo/Go/plugin"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(helpBot{priority: 998})
}

type helpBot struct {
	priority int //[0~1000)
}

func (h helpBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	if !req.IsAtMe {
		return false, true
	}
	return true, false
}

func (h helpBot) Process(req *plugin.Request) []*plugin.Result {
	//var cmd,cmd2,cmd3 string
	//flagSet:=flag.NewFlagSet("修玛吉亚", 0)
	//var buf bytes.Buffer
	//flagSet.SetOutput(&buf)
	//flagSet.
	//flagSet.StringVar(&cmd,"公告","","获取所有现有公告")
	//flagSet.StringVar(&cmd2,"建造","","支持的建筑物:水祭坛,觉醒之岚树")
	//flagSet.StringVar(&cmd3,"抽卡","","支持的抽卡方式:单抽,十连,百连")
	////flagSet.StringVar(&cmd4,"招募","","招募车队")
	//args:=strings.Split(req.Content," ")
	//flagSet.Parse(args[1:])
	//
	//fmt.Println(cmd,cmd2,cmd3)
	//flagSet.Usage()
	res := ""
	switch req.Content {
	case "帮助":
		fallthrough
	case "help":
		res += "不需要at我就能直接触发的有:单抽 十连 百连 千连 万连可以触发抽卡逻辑\n"
		res += "输入表情包的名字可以获取对应表情一张\n"
		res += "招募xx缺x可以触发招募逻辑\n"
		//res += "呼唤人气群友的名字可以获得灵魂拷问一次\n"
		res += "任何时候都可以得到我的复读,获得召唤卷若干\n"
		res += "需要at我触发的有:@我 建造xxx 暂时只支持水祭坛/觉醒之岚树/金币矿山\n"
		//res += "@我 公告 获取最新信息\n"
		res += "@我 图鉴 查看图鉴\n"
		res += "@我 成就/统计 查看统计信息和成就\n"
		//res += "@我  查看统计信息和成就\n"
		res += "@我 转生 进行转生,转生一次永久提高0.5%虹率\n"
	case "碎碎念":
		res += "本人特别喜欢龙约这款游戏,时常在小群约车比较难招募到人,所以本意想做一个群车招募机器人.后来群友建议做个抽卡机器人," +
			"进化的过程略过不提,最后做到今天这个样子.我希望给机器人的定位是活跃群气氛,产生一些戏剧效果,是一个平时聊天灌水的添头." +
			"我会不断的修改机器人的赠送逻辑以及其他功能以符合机器人的最初定位,不希望机器人影响正常群里的聊天~" +
			//"然后为了让大家玩的开心做了存档和收集功能,这些数据我尽量备份但是因为精力/技术都有限," +
			//"很多地方做的比较粗糙,数据偶尔丢失以及使用人数多了之后的bug都在所难免.欢迎各位github上反馈给我~" +
			"很多地方做的比较粗糙,各位海涵.有问题欢迎各位github上反馈给我~" +
			//"如果需要添加到自己的小群,QQ私聊我一下就好了~" +
			"github上搜DragaliaLostQQBot即可 欢迎大家提issues/建议/创意和贡献代码"
	case "github":
		res += "https://github.com/yucendulang/DragaliaLostQQBot"
	default:
		//if req.Conf.IsIntroOn {
		res += "@我 公告 获取最新信息;帮助 查看各种指令;碎碎念 一些无聊的话"
		//} else {
		//	res += "@我 公告 获取最新信息;帮助 查看各种指令"
		//}
	}

	return []*plugin.Result{{
		Content: res,
	}}
}

func (h helpBot) Priority() int {
	return h.priority
}
