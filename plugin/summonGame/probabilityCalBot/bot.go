package probabilityCalBot

import (
	"bytes"
	"flag"
	"fmt"
	"iotqq-plugins-demo/Go/plugin"
	"strings"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&collectorBot{5})
}

type collectorBot struct {
	priority int //[0~1000)
}

func (c *collectorBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	f := plugin.NewCommonPrefixTriggerFunc("概率计算")
	return f(req)
}

func (c *collectorBot) Process(req *plugin.Request) []*plugin.Result {
	args := strings.Split(req.Content, " ")
	flagSet := flag.NewFlagSet("概率计算", 0)
	var goalCard = flagSet.String("g", "", "goal card,split by comma")
	var drawNum = flagSet.Int("d", 0, "draw time")
	var cardPoolIndex = flagSet.Int("p", 0, "pool index")
	var buf bytes.Buffer
	flagSet.SetOutput(&buf)
	flagSet.Parse(args[1:])

	if (*goalCard) == "" {
		flagSet.Usage()
		return []*plugin.Result{{Content: buf.String()}}
	}

	fmt.Println(*goalCard, *drawNum, *cardPoolIndex)
	content, f, err := SimParse(*goalCard, *drawNum, *cardPoolIndex)
	if err != nil {
		return []*plugin.Result{{Content: err.Error()}}
	} else {
		return []*plugin.Result{{Content: content, DelayFunc: f}}
	}
}

func (c *collectorBot) Priority() int {
	return c.priority
}
