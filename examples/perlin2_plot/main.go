// examples/perlin2_plot/main.go
//
// Visualize fuzzl.Perlin2 over a grid using gg.
// Usage:
//   go run ./examples/perlin2_plot                    # grayscale
//   go run ./examples/perlin2_plot -heatmap           # heatmap
//   go run ./examples/perlin2_plot -scale 0.02        # zoom in
//   go run ./examples/perlin2_plot -scale 0.08        # zoom out
//
// Output: testdata/perlin2.png

package main

import (
	"flag"
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/mdhender/fuzzl"
)

func main() {
	heatmap := flag.Bool("heatmap", false, "render using a heatmap instead of grayscale")
	scale := flag.Float64("scale", 0.04, "world units per pixel (smaller = zoom in)")
	flag.Parse()

	const (
		width  = 512
		height = 512
	)
	dc := gg.NewContext(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			px := (float64(x) - float64(width)/2) * (*scale)
			py := (float64(y) - float64(height)/2) * (*scale)

			v := fuzzl.Perlin2(fuzzl.Vec2{px, py}) // ~[0,1]

			if *heatmap {
				r, g, b := hsvToRGB(v*360.0, 1.0, 1.0)
				dc.SetRGB(r, g, b)
			} else {
				dc.SetRGB(v, v, v)
			}
			dc.SetPixel(x, y)
		}
	}

	if err := dc.SavePNG("testdata/perlin2.png"); err != nil {
		log.Fatal(err)
	}
}

// hsvToRGB converts HSV to RGB (all in [0,1], hue in degrees).
func hsvToRGB(h, s, v float64) (r, g, b float64) {
	c := v * s
	hh := h / 60.0
	x := c * (1 - math.Abs(math.Mod(hh, 2)-1))
	var rr, gg, bb float64

	switch {
	case hh >= 0 && hh < 1:
		rr, gg, bb = c, x, 0
	case hh >= 1 && hh < 2:
		rr, gg, bb = x, c, 0
	case hh >= 2 && hh < 3:
		rr, gg, bb = 0, c, x
	case hh >= 3 && hh < 4:
		rr, gg, bb = 0, x, c
	case hh >= 4 && hh < 5:
		rr, gg, bb = x, 0, c
	case hh >= 5 && hh < 6:
		rr, gg, bb = c, 0, x
	}
	m := v - c
	return rr + m, gg + m, bb + m
}
