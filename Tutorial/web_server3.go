package main

import (
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"strconv"
	"fmt"
	"io"
)

var pallete = []color.Color{color.White, color.RGBA{0, 255, 0, 255}}
const (
	whiteIndex = 0 // first color in pallete
	blackIndex = 1 // next color in pallete
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cyclesStr := r.FormValue("cycles")
		if cyclesStr == "" {
			fmt.Printf("Param cycles is required.\n")
			return
		}
		cycles, err := strconv.Atoi(cyclesStr)
		if err != nil {
			fmt.Fprintf(w, "Bad cycles value: %s\n", err)
		}
		lissajous(cycles, w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajous(cycles int, out io.Writer) {
	const (
		res    	= 0.001 // angular resolution
		size   	= 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay   = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscilattor
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	counter := 0
	counterRed := 255
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		counter += 5
		counterRed -= 10
		var pallete = []color.Color{color.White, color.RGBA{uint8(counterRed), uint8(counter), 0, uint8(counter)}}
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		} 
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
