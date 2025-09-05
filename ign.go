package fuzzl

import "math"

// IGN returns a deterministic pseudo-random value in [0,1) for the given pixel.
func IGN(pixelX, pixelY int) float64 {
	x := float64(pixelX)
	y := float64(pixelY)
	t := math.Mod(0.06711056*x+0.00583715*y, 1.0)
	return math.Mod(52.9829189*t, 1.0)
}

// IGNByFrame is a frame-varying variant of IGN.
// The frame index is wrapped every 64 steps to avoid numerical drift.
// Results are still in [0,1).
func IGNByFrame(pixelX, pixelY, frame int) float64 {
	const FrameShift = 5.588238

	// Wrap frame to [0,63]
	f := frame % 64

	x := float64(pixelX) + FrameShift*float64(f)
	y := float64(pixelY) + FrameShift*float64(f)

	t := math.Mod(0.06711056*x+0.00583715*y, 1.0)
	return math.Mod(52.9829189*t, 1.0)
}
