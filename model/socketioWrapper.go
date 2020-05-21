package model

import (
	gosocketio "github.com/graarh/golang-socketio"
	"log"
	"time"
)

func SendJoin(c *gosocketio.Client, qq string) {
	log.Println("获取QQ号连接")
	result, err := c.Ack("GetWebConn", qq, time.Second*5)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("emit", result)
	}
}
