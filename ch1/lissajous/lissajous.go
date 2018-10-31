// Lissajous generates animated GIF from
// random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // First palette color
	blackIndex = 1 // Next palette color
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles    = 5
		res       = 0.001 //Sin() function resolution
		size      = 100
		nframes   = 64 // Animation frames count
		delay     = 8  // Delay between frames (multiplicated by 10ms)
		loopCount = 0  // Loop forever
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: loopCount}
	phase := 0.0 // Phases difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
