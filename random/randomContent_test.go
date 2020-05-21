package random

import (
	"fmt"
	"testing"
)

func Test_randomGetSuffix(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(RandomGetSuffix())
		})
	}
}
