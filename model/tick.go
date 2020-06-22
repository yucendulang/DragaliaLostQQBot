package model

import (
	"fmt"
	"iotqq-plugins-demo/Go/building"
	"iotqq-plugins-demo/Go/common"
	"iotqq-plugins-demo/Go/userData"
	"time"
)

func StartTick() {
	go func() {
		for range time.Tick(time.Minute * 1) {
			userData.UserRange(func(key, value interface{}) bool {
				userInfo := value.(*userData.User)
				num := building.GetBuildEffect(userInfo.BuildIndex).VolunterMineProduct
				if num == 0 {
					return true
				}
				if time.Since(userInfo.LastVolunterGetTime) > common.VolunterMineProductPeriod {
					userInfo.SummonCardNum += num
					userInfo.LastVolunterGetTime = time.Now()
					userInfo.Static.VRTPeriod = 0
					userData.SaveUserByUDID(userInfo.Udid)
					fmt.Println(userInfo.Udid, num, userInfo.LastVolunterGetTime)
				}
				return true
			})
		}
	}()
}
