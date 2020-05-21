package model

import "encoding/json"

type QQMsg struct {
	Content string
	UserID  []int64
}

func NewQQMsg(msg string) (*QQMsg, error) {
	res := &QQMsg{}
	err := json.Unmarshal([]byte(msg), res)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
