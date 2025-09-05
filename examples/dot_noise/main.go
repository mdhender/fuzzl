// dotnoise_plot.go
//
// Example program to visualize fuzzl.DotNoise across a 2D plane
// using the gg (github.com/fogleman/gg) package.
package main

import (
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/mdhender/fuzzl"
)

func main() {
	const (
		width  = 512
		height = 512
		scale  = 0.05 // controls zoom
		z      = 0.0  // fixed z slice
	)

	dc := gg.NewContext(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Map pixel coordinates to noise coordinates
			px := (float64(x) - width/2) * scale
			py := (float64(y) - height/2) * scale
			p := fuzzl.Vec3{px, py, z}

			v := fuzzl.DotNoise(p)

			// Normalize from [-3,3] -> [0,1]
			norm := (v + 3.0) / 6.0
			norm = math.Max(0, math.Min(1, norm)) // clamp

			// Draw pixel
			gray := norm
			dc.SetRGB(gray, gray, gray)
			dc.SetPixel(x, y)
		}
	}

	if err := dc.SavePNG("testdata/dotnoise.png"); err != nil {
		log.Fatal(err)
	}
}
