package userData

import (
	"fmt"
	"testing"
)

func TestUserDataSave(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserDataSave()
		})
	}
}

func TestUserDataLoad(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "basic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserDataLoad()
			UserMap.Range(func(key, value interface{}) bool {
				fmt.Printf("key:%v,value:%v\n", key, value)
				return true
			})
		})
	}
}
