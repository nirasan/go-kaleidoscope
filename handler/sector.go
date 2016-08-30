package handler

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
)

func Sector(w http.ResponseWriter, r *http.Request) {
	var (
		width           = 300
		height          = 300
		centerX         = 150
		centerY         = 150
		sectorR float64 = 100
		sectorD float64 = 30
	)
	colors := []color.NRGBA{
		color.NRGBA{0, 0, 0, 255},
		color.NRGBA{128, 128, 128, 255},
		color.NRGBA{255, 255, 255, 255},
	}
	// 縦縞
	img := image.NewPaletted(image.Rect(0, 0, width, height), color.Palette{colors[0], colors[1], colors[2]})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Pix[y*img.Stride+x] = uint8((x / 5) % 3)
		}
	}
	// 扇形に画像を切り抜いて回転させる
	img2 := image.NewPaletted(image.Rect(0, 0, width, height), color.Palette{colors[2], colors[1], colors[0]})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if contain(sectorR, sectorD, centerX, centerY, centerX, height, x, y) {
				for _, deg := range []float64{60, -60, 120, -120, 180} {
					rad := deg * math.Pi / 180
					x1, y1 := float64(x-centerX), float64(y-centerY)
					x2 := x1*math.Cos(rad) - y1*math.Sin(rad)
					y2 := x1*math.Sin(rad) + y1*math.Cos(rad)
					x3, y3 := int(math.Floor(x2+.5))+centerX, int(math.Floor(y2+.5))+centerY
					img2.Pix[y3*img.Stride+x3] = img.Pix[y*img.Stride+x]
				}
				img2.Pix[y*img.Stride+x] = img.Pix[y*img.Stride+x]
			}
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img2)
}

// contain は点A(x1,y1)を中心、点Aと点Cの線を中心、d*2を角度、rを半径とする扇形に点Cが含まれるかを判定する
func contain(r, d float64, x1, y1, x2, y2, x3, y3 int) bool {
	xa, ya := x2-x1, y2-y1
	xb, yb := x3-x1, y3-y1
	ac := float64(xa*xb + ya*yb)
	an := math.Pow(float64(xa), 2) + math.Pow(float64(ya), 2)
	bn := math.Pow(float64(xb), 2) + math.Pow(float64(yb), 2)
	rad := math.Acos(ac / math.Sqrt(an*bn))
	deg := rad * 180 / math.Pi
	dist := math.Sqrt(math.Pow(float64(xb), 2) + math.Pow(float64(yb), 2))
	if deg <= d && dist < r {
		return true
	}
	return false
}
