package probabilityCalBot

import (
	"fmt"
	"iotqq-plugins-demo/Go/plugin"
	"testing"
	"time"
)

func Test_collectorBot_Process(t *testing.T) {
	type fields struct {
		priority int
	}
	type args struct {
		req *plugin.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		//{"basic", fields{}, args{req: &plugin.Request{
		//	Content: "概率计算 -g 莉莉",
		//}}},
		{"basicAllIn", fields{}, args{req: &plugin.Request{
			Content: "概率计算 -g 莉莉 -d 1000",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &collectorBot{
				priority: tt.fields.priority,
			}
			if got := c.Process(tt.args.req); got[0].DelayFunc != nil {
				result := got[0].DelayFunc()
				fmt.Println(result)
				time.Sleep(time.Second * 1)
			}
		})
	}
}
