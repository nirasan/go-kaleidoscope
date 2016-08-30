package handler

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
)

// Sector のピクセル抜けの修正
// http://homepage2.nifty.com/tsugu/sotuken/rotation/#bi_liner
func Sector2(w http.ResponseWriter, r *http.Request) {
	var (
		width           = 300
		height          = 300
		centerX         = 150
		centerY         = 150
		sectorR float64 = 100
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
	// 扇形の回転軌道上にいる場合は元の座標から画素をとってくる
	img2 := image.NewPaletted(image.Rect(0, 0, width, height), color.Palette{colors[2], colors[1], colors[0]})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			degree, ok := getDegree(sectorR, centerX, centerY, centerX, height, x, y)
			if ok {
				for _, deg := range []float64{0, 60, 120, 180} {
					if degree < deg-30 || deg+30 <= degree {
						continue
					}
					// 180 度以上は逆回転させる
					if x < centerX {
						deg *= -1
					}
					rad := deg * math.Pi / 180
					x1 := float64(x-centerX)*math.Cos(rad) - float64(y-centerY)*math.Sin(rad) + float64(centerX)
					y1 := float64(x-centerX)*math.Sin(rad) + float64(y-centerY)*math.Cos(rad) + float64(centerY)
					x2, y2 := int(math.Floor(x1+.5)), int(math.Floor(y1+.5))
					if len(img.Pix) > y2*img.Stride+x2 {
						img2.Pix[y*img.Stride+x] = img.Pix[y2*img.Stride+x2]
					}
				}
			}
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img2)
}

// x1, y1 と x2, y2 を中心線とする扇形の回転上にあったら中心線との角度を返す
func getDegree(r float64, x1, y1, x2, y2, x3, y3 int) (float64, bool) {
	xa, ya := x2-x1, y2-y1
	xb, yb := x3-x1, y3-y1
	ac := float64(xa*xb + ya*yb)
	an := math.Pow(float64(xa), 2) + math.Pow(float64(ya), 2)
	bn := math.Pow(float64(xb), 2) + math.Pow(float64(yb), 2)
	rad := math.Acos(ac / math.Sqrt(an*bn))
	deg := rad * 180 / math.Pi
	dist := math.Sqrt(math.Pow(float64(xb), 2) + math.Pow(float64(yb), 2))
	if dist < r {
		return deg, true
	}
	return 0, false
}
