package kaleidoscope

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type Position struct {
	X, Y int
}

type Circle struct {
	*Position
	R     int
	Color color.RGBA
}

func (c *Circle) Contain(x, y int) bool {
	r := math.Sqrt(math.Pow(float64(x-c.Position.X), 2) + math.Pow(float64(y-c.Position.Y), 2))
	return r <= float64(c.R)
}

type Rect struct {
	*Position
	W, H  int
	Color color.RGBA
}

func (r *Rect) Contain(x, y int) bool {
	x1, y1 := r.Position.X-r.W/2, r.Position.Y-r.H/2
	x2, y2 := x1+r.W, y1+r.H
	return x1 <= x && x <= x2 && y1 <= y && y <= y2
}

var (
	defaultWidth           = 300
	defaultHeight          = 300
	defaultCenterX         = 150
	defaultCenterY         = 150
	defaultSectorR float64 = 145
	defaultNum             = 30
)

func CreateImage(w, n int) *image.RGBA {

	rand.Seed(time.Now().UnixNano())

	width := defaultWidth
	height := defaultHeight
	centerX := defaultCenterX
	centerY := defaultCenterY
	sectorR := defaultSectorR

	// 画像サイズ
	if w != 0 {
		width = w
		height = width
		centerX = width / 2
		centerY = height / 2
		sectorR = float64(centerX) * 0.9
	}

	// 生成要素数
	num := defaultNum
	if n != 0 {
		num = n
	}

	rectList := []Rect{}
	circleList := []Circle{}

	// 丸と四角の生成
	for i := 0; i < num; i++ {
		rect := Rect{
			Position: &Position{X: rand.Intn(width), Y: rand.Intn(height)},
			Color:    randomColor(),
			W:        rand.Intn(50),
			H:        rand.Intn(50),
		}
		rectList = append(rectList, rect)
		circle := Circle{
			Position: &Position{X: rand.Intn(width), Y: rand.Intn(height)},
			Color:    randomColor(),
			R:        rand.Intn(50),
		}
		circleList = append(circleList, circle)
	}

	// 丸と四角の描画
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.RGBA{255, 255, 255, 255}
			skip := false
			for _, circle := range circleList {
				if !skip && circle.Contain(x, y) {
					c = circle.Color
					skip = true
				}
			}
			for _, rect := range rectList {
				if !skip && rect.Contain(x, y) {
					c = rect.Color
					skip = true
				}
			}
			img.SetRGBA(x, y, c)
		}
	}

	// 扇形の範囲をコピーして回転
	// 扇形の回転軌道上にいる場合は元の座標から画素をとってくる
	img2 := image.NewRGBA(image.Rect(0, 0, width, height))
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
						img2.SetRGBA(x, y, img.RGBAAt(x2, y2))
					}
				}
			}
		}
	}

	return img2
}

func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
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
