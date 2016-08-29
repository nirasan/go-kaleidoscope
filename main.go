package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	colors := []color.NRGBA{
		color.NRGBA{0, 0, 0, 255},
		color.NRGBA{128, 128, 128, 255},
		color.NRGBA{255, 255, 255, 255},
	}
	img := image.NewPaletted(image.Rect(0, 0, 100, 100), color.Palette{colors[0], colors[1], colors[2]})
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Pix[y*img.Stride+x] = uint8(rand.Intn(3))
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
