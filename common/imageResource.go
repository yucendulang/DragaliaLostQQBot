package common

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
)

var Font *truetype.Font

func init() {
	fontSourceBytes, _ := ioutil.ReadFile("./asset/msyhbd.ttc")
	Font, _ = truetype.Parse(fontSourceBytes)
}
