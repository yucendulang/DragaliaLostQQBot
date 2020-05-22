package util

import (
	"image"
	"image/color"
)

func Clear(blackBar *image.RGBA, c color.RGBA) {
	dx, dy := blackBar.Rect.Dx(), blackBar.Rect.Dy()
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			blackBar.SetRGBA(i, j, c)
		}
	}
}
