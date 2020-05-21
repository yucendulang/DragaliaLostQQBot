package util

import "testing"

func TestFixName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{name: "aaa2345-45-345"}, "aaa"},
		{"basic", args{name: "aaa2345-45-3450"}, "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixName(tt.args.name); got != tt.want {
				t.Errorf("FixName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixSentense(t *testing.T) {
	type args struct {
		sentense string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"basic", args{sentense: "我也要呢!!!!!"}, "我也要"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixSentense(tt.args.sentense); got != tt.want {
				t.Errorf("FixSentense() = %v, want %v", got, tt.want)
			}
		})
	}
}
