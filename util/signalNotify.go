package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SignalNotify() {
	fmt.Println("enter SignalNotify")
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGINT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT:
				fmt.Println("尝试存档内存的数据", s)
				ramVarMap.Range(func(key, value interface{}) bool {
					f, err := os.Create(ramVarPath + key.(string))
					defer f.Close()

					info, _ := f.Stat()
					fmt.Println("尝试存档", info.Name())
					if err != nil {
						fmt.Println(err.Error())
					} else {
						b, err := json.Marshal(value)
						_, err = f.Write(b)
						if err != nil {
							fmt.Println(err.Error())
						}
					}
					return true
				})
				fmt.Println("退出程序")
				os.Exit(0)
			default:
				fmt.Println("other", s)
			}
		}
	}()
}
