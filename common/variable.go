package common

import (
	"strconv"
	"time"
)

var FiveStarCharacterNum, FourStarCharacterNum, ThreeStarCharacterNum int
var FiveStarDragonNum, FourStarDragonNum, ThreeStarDragonNum int
var GachaPoolNum, NoGachaPoolNum int

var VolunterMineProductPeriod = time.Hour * 6
var BaseSSRProbality = 40

var QQ = "2834323101"
var QQInt int64

func init() {
	i, _ := strconv.Atoi(QQ)
	QQInt = int64(i)
}

//var SelfQQ int64
