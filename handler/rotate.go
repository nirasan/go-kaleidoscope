package handler

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
)

func Rotate(w http.ResponseWriter, r *http.Request) {
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
	for x := 50; x < 100; x++ {
		for y := 50; y < 100; y++ {
			for _, deg := range []float64{90, -90} {
				rad := deg * math.Pi / 180
				// 中心点の分移動
				x1, y1 := float64(x-50), float64(y-50)
				// 座標の回転
				x2 := x1*math.Cos(rad) - y1*math.Sin(rad)
				y2 := x1*math.Sin(rad) + y1*math.Cos(rad)
				// 中心点の分もどす
				x3, y3 := int(x2+50), int(y2+50)
				img.Pix[y3*img.Stride+x3] = img.Pix[y*img.Stride+x]
			}
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
