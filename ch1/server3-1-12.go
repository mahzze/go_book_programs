package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

/*
This is exercise 1.12, which is "Exercise 1.12: Modify the Lissajous server to
read parameter values from the URL.For example, you might arrange it so that a
URL like http://localhost:8000/?cycles=20 sets the number of cycles to 20
instead of the default 5. Use the strconv.Atoi function to convert the string
parameter into an integer"
*/
var cycles = 5.0 // integers from the original program must be changed to floats to avoid float truncation
var res = 0.001
var size = 100
var nframes = 64
var delay = 8

func main() {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		parameters, _ := url.ParseQuery(request.URL.RawQuery)

		if len(parameters["cycles"]) != 0 {
			cyclesInput, err := strconv.ParseFloat(parameters["cycles"][0], 64)
			checkFloat(err, &cyclesInput, &cycles)
		}
		if len(parameters["res"]) != 0 {
			resInput, err := strconv.ParseFloat(parameters["res"][0], 64)
			checkFloat(err, &resInput, &res)
		}
		if len(parameters["size"]) != 0 {
			sizeInput, err := strconv.Atoi(parameters["size"][0])
			checkInt(err, &sizeInput, &size)
		}
		if len(parameters["nframes"]) != 0 {
			nframesInput, err := strconv.Atoi(parameters["nframes"][0])
			checkInt(err, &nframesInput, &nframes)
		}
		if len(parameters["delay"]) != 0 {
			delayInput, err := strconv.Atoi(parameters["delay"][0])
			checkInt(err, &delayInput, &delay)
		}
		lissajous(writer, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func checkFloat(err error, input *float64, variable *float64) {
	if err != nil {
		panic(err)
	} else {
		*variable = *input
	}
}

func checkInt(err error, input *int, variable *int) {
	if err != nil {
		panic(err)
	} else {
		*variable = *input
	}
}

const (
	color_index     = 1
	black_bkg_index = 0
)

func lissajous(out io.Writer, cycles float64, res float64, size int, nframes int, delay int) {

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, []color.Color{
			color.Black,
			color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 255},
		})
		for t := 0.0; t < cycles*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+.5), size+int(y*float64(size)+.5), color_index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
