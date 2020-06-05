package common

import "time"

type BuildRecord struct {
	Index, Level int
}

type AchievementRecord struct {
	Index           int
	AchievementTime time.Time
}
