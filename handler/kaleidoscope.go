package handler

import (
	"github.com/nirasan/go-kaleidoscope/kaleidoscope"
	"image/png"
	"net/http"
	"strconv"
)

func Kaleidoscope(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	// 画像サイズ
	width := 0
	if ws := q.Get("w"); ws != "" {
		var err error
		width, err = strconv.Atoi(ws)
		if err != nil {
			http.Error(w, "Invalid w parameter, must be an integer", http.StatusBadRequest)
			return
		}
	}

	// 生成要素数
	num := 0
	if ns := q.Get("n"); ns != "" {
		var err error
		num, err = strconv.Atoi(ns)
		if err != nil {
			http.Error(w, "Invalid n parameter, must be an integer", http.StatusBadRequest)
			return
		}
	}

	img := kaleidoscope.CreateImage(width, num)

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
