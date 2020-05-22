package plugin

import (
	"fmt"
	"iotqq-plugins-demo/Go/model"
	"iotqq-plugins-demo/Go/util"
	"regexp"
	"sort"
	"strings"
)

type Result struct {
	Content   string
	PicUrl    string
	DelayFunc func() string
}
type Request struct {
	Udid      int64
	Content   string
	NickName  string
	IsAtMe    bool
	ExtraInfo interface{}
	GroupPics []model.GroupPic
}

type Interface interface {
	IsTrigger(req *Request) (res bool, vNext bool)
	Process(req *Request) *Result
	Priority() int
}

type Factory struct {
	seftUDID   int64
	interfaces []Interface
}

type ifs []Interface

func (is ifs) Len() int {
	return len(is)
}

func (is ifs) Less(i, j int) bool {
	return is[i].Priority() < is[j].Priority()
}

func (is ifs) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

var FactoryInstance Factory

func (f *Factory) SetOptions(selfQQ int64) {
	FactoryInstance.seftUDID = selfQQ
}

func (f *Factory) RegisterPlugin(i Interface) {
	fmt.Println(i, "register success")
	(*f).interfaces = append((*f).interfaces, i)
	sort.Sort(ifs((*f).interfaces))
}

var AtMsgRegex = regexp.MustCompile("^@.*? (.*)$")

func (f Factory) Run(data model.Data) {
	req := &Request{Udid: data.FromUserID, NickName: util.FixName(data.FromNickName)}
	if strings.HasPrefix(data.Content, "{") {
		if msg, err := model.NewQQMsg(data.Content); err != nil {
			req.Content = data.Content
		} else {
			//json 解析成功
			if util.Int64Contain(f.seftUDID, msg.UserID) {
				find := AtMsgRegex.FindStringSubmatch(msg.Content)
				if len(find) > 0 {
					req.Content = find[1]
				} else {
					req.Content = msg.Content
				}
				req.IsAtMe = true
			} else {
				req.Content = msg.Content
			}
			req.GroupPics = msg.GroupPic
		}
	} else {
		req.Content = data.Content
	}

	req.Content = strings.TrimSpace(req.Content)

	for i := 0; i < len(f.interfaces); i++ {
		p := f.interfaces[i]
		r, vNext := p.IsTrigger(req)
		if r {
			res := p.Process(req)
			if res.PicUrl != "" {
				model.SendPic(data.FromGroupID, 2, res.Content, res.PicUrl)
			} else if res.Content != "" {
				model.Send(data.FromGroupID, 2, res.Content)
			}
			if res.DelayFunc != nil {
				go func() {
					fmt.Printf("enter DelayFunc")
					outStr := res.DelayFunc()
					model.Send(data.FromGroupID, 2, outStr)
				}()
			}
		}
		if !vNext {
			return
		}
	}
}
