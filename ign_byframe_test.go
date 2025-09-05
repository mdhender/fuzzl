package fuzzl

import "testing"

func TestIGNByFrameRange(t *testing.T) {
	for frame := 0; frame < 128; frame++ { // test wrapping
		for _, pt := range [][2]int{{0, 0}, {5, 7}, {-3, 2}, {128, 256}} {
			v := IGNByFrame(pt[0], pt[1], frame)
			if v < 0 || v >= 1 {
				t.Fatalf("IGNByFrame(%d,%d,%d) out of range: %v",
					pt[0], pt[1], frame, v)
			}
		}
	}
}
