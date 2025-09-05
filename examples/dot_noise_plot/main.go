// Example program to visualize fuzzl.DotNoise across a 2D plane
// using the gg (github.com/fogleman/gg) package.
//
// Usage:
//
//  go run ./examples/dot_noise_plot                      # grayscale (default)
//  go run ./examples/dot_noise_plot -heatmap             # heatmap coloring
//  go run ./examples/dot_noise_plot -z 5.0               # grayscale at z=5.0
//  go run ./examples/dot_noise_plot -z 2.5 -heatmap      # heatmap at z=2.5
//  go run ./examples/dot_noise_plot -scale 0.1           # zoom out (larger features)
//  go run ./examples/dot_noise_plot -z 2.5 -scale 0.02   # zoom in slice at z=2.5

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
	z := flag.Float64("z", 0.0, "z-slice coordinate for the noise field")
	scale := flag.Float64("scale", 0.05, "zoom level (larger = zoom out, smaller = zoom in)")
	flag.Parse()

	const (
		width  = 512
		height = 512
	)

	dc := gg.NewContext(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Map pixel coordinates to noise coordinates
			px := (float64(x) - width/2) * (*scale)
			py := (float64(y) - height/2) * (*scale)
			p := fuzzl.Vec3{px, py, *z}

			v := fuzzl.DotNoise(p)

			// Normalize from [-3,3] -> [0,1]
			norm := (v + 3.0) / 6.0
			norm = math.Max(0, math.Min(1, norm)) // clamp

			if *heatmap {
				r, g, b := hsvToRGB(norm*360.0, 1.0, 1.0)
				dc.SetRGB(r, g, b)
			} else {
				dc.SetRGB(norm, norm, norm)
			}
			dc.SetPixel(x, y)
		}
	}

	if err := dc.SavePNG("testdata/dot_noise.png"); err != nil {
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
