package model

import (
	"time"
)

func Periodlycall(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
		//fmt.Println(x)
	}
}
