package stickerBot

import "testing"

func TestCacheStickerFile(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{"basic", args{url: stickerMap["不行"].url[0]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CacheStickerFile(tt.args.url)
		})
	}
}
