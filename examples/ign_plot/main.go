// examples/ign_plot/main.go
//
// Example program to visualize fuzzl.IGN or fuzzl.IGNByFrame
// using gg (github.com/fogleman/gg).
//
// Usage:
//
//	go run ./examples/ign_plot                   # grayscale IGN
//	go run ./examples/ign_plot -frame 5          # IGNByFrame at frame 5
//	go run ./examples/ign_plot -frame 10 -heatmap
//
// The generated image is saved in testdata/ign.png.
package main

import (
	"flag"
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/mdhender/fuzzl"
)

func main() {
	frame := flag.Int("frame", -1, "frame index (use -1 for static IGN)")
	heatmap := flag.Bool("heatmap", false, "render using a heatmap instead of grayscale")
	flag.Parse()

	const (
		width  = 512
		height = 512
	)

	dc := gg.NewContext(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var v float64
			if *frame >= 0 {
				v = fuzzl.IGNByFrame(x, y, *frame)
			} else {
				v = fuzzl.IGN(x, y)
			}

			if *heatmap {
				// map [0,1) to HSV → RGB
				r, g, b := hsvToRGB(v*360.0, 1.0, 1.0)
				dc.SetRGB(r, g, b)
			} else {
				dc.SetRGB(v, v, v)
			}
			dc.SetPixel(x, y)
		}
	}

	if err := dc.SavePNG("testdata/ign.png"); err != nil {
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
