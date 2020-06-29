package model

import (
	"fmt"
	gosocketio "github.com/graarh/golang-socketio"
	"time"
)

func SendJoin(c *gosocketio.Client, qq string) bool {
	fmt.Println("获取QQ号连接")
	result, err := c.Ack("GetWebConn", qq, time.Second*5)
	if err != nil {
		fmt.Println("GetWebConn返回错误:", err)
		return false
	} else if result != "\"OK\"" {
		fmt.Println("result返回不是OK,尝试重新链接中:", result)
		return false
	} else {
		fmt.Println("emit", result)
		return true
	}
}
