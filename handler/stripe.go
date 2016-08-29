package handler

import (
	"image"
	"image/color"
	"image/png"
	"net/http"
)

func Stripe(w http.ResponseWriter, r *http.Request) {
	colors := []color.NRGBA{
		color.NRGBA{0, 0, 0, 255},
		color.NRGBA{128, 128, 128, 255},
		color.NRGBA{255, 255, 255, 255},
	}
	// 縦縞
	img := image.NewPaletted(image.Rect(0, 0, 100, 100), color.Palette{colors[0], colors[1], colors[2]})
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Pix[y*img.Stride+x] = uint8((x / 5) % 3)
		}
	}
	// 90度回転
	img2 := image.NewPaletted(image.Rect(0, 0, 100, 100), color.Palette{colors[0], colors[1], colors[2]})
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img2.Pix[y*img2.Stride+x] = img.Pix[x*img.Stride+y]
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img2)
}
