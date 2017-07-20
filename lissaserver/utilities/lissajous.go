package utilities

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x42, 0xa1, 0xf4, 1},
	color.RGBA{0x42, 0x89, 0xf4, 1},
	color.RGBA{0x42, 0x7d, 0xf4, 1},
	color.RGBA{0x42, 0x68, 0xf4, 1},
	color.RGBA{0x42, 0x56, 0xf4, 1},
	color.RGBA{0x42, 0x56, 0xf4, 1},
	color.RGBA{0x42, 0x68, 0xf4, 1},
	color.RGBA{0x42, 0x7d, 0xf4, 1},
	color.RGBA{0x42, 0x89, 0xf4, 1},
	color.RGBA{0x42, 0xa1, 0xf4, 1}}

const (
	whiteIndex = 0
	blackIndex = 1
)

// Lissajous write Lissajous to stream
func Lissajous(out http.ResponseWriter, cycles float64) {
	const (
		res     = 0.001
		size    = 200
		nframes = 200
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8((i % (len(palette) - 1)) + 1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
