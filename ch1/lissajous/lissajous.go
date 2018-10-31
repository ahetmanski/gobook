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

var palette = color.Palette{
	color.Black,
	color.RGBA{230, 25, 75, 0xFF},
	color.RGBA{60, 180, 75, 0xFF},
	color.RGBA{255, 225, 25, 0xFF},
	color.RGBA{0, 130, 200, 0xFF},
	color.RGBA{245, 130, 48, 0xFF},
	color.RGBA{145, 30, 180, 0xFF},
	color.RGBA{70, 240, 240, 0xFF},
	color.RGBA{240, 50, 230, 0xFF},
	color.RGBA{210, 245, 60, 0xFF},
	color.RGBA{250, 190, 190, 0xFF},
	color.RGBA{0, 128, 128, 0xFF},
	color.RGBA{230, 190, 255, 0xFF},
	color.RGBA{170, 110, 40, 0xFF},
	color.RGBA{255, 250, 200, 0xFF},
	color.RGBA{128, 0, 0, 0xFF},
	color.RGBA{170, 255, 195, 0xFF},
	color.RGBA{128, 128, 0, 0xFF},
	color.RGBA{255, 215, 180, 0xFF},
	color.RGBA{0, 0, 128, 0xFF},
	color.RGBA{128, 128, 128, 0xFF},
}

const (
	backgroundIndex = 0
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
		lineColorIndex := rand.Intn(len(palette))
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(lineColorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
