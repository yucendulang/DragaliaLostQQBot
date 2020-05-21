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

func (h helpBot) Process(req *plugin.Request) *plugin.Result {
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
	res := "我没有理解你想做什么,现在我能提供以下服务\n"
	res += "不需要at我就能直接触发的有:单抽 十连 百连 可以触发抽卡逻辑\n"
	res += "输入表情包的名字可以获取对应表情一张\n"
	res += "招募xx缺x可以触发招募逻辑\n"
	res += "呼唤人气群友的名字可以获得灵魂拷问一次\n"
	res += "任何时候都可以得到修玛吉亚的复读,获得召唤卷若干\n"
	res += "需要at我触发的有:@我 建造xxx 暂时只支持水祭坛/觉醒之岚树\n"
	res += "@我 公告 获取最新信息\n"
	return &plugin.Result{
		Content: res,
		//Content: buf.String(),
	}
}

func (h helpBot) Priority() int {
	return h.priority
}
