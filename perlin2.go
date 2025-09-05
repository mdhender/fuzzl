// 2D Perlin noise (shader-style) based on Xor’s explanation at GM Shaders.
// Perlin noise interpolates between *gradient* directions at cell corners
// using quintic smoothing for natural results.
//
// GM Shaders is a reader-supported publication.
// Consider supporting: https://gmshaders.com

package fuzzl

import "math"

// hash2 returns a pseudo-random Vec2 in [0,1)^2 for a given 2D point p.
// GLSL form: fract(sin(p * mat2(0.129898,0.78233, 0.81314,0.15926)) * 43758.5453)
func hash2(p Vec2) Vec2 {
	// Expanded: p.x * vec2(0.129898, 0.81314) + p.y * vec2(0.78233, 0.15926)
	x := p[0]*0.129898 + p[1]*0.78233
	y := p[0]*0.81314 + p[1]*0.15926
	sx := math.Sin(x) * 43758.5453
	sy := math.Sin(y) * 43758.5453
	return Vec2{fract(sx), fract(sy)}
}

// hash2Norm returns a *unit* vector derived from hash2(p).
// GLSL: normalize(hash2(p) - 0.5)
func hash2Norm(p Vec2) Vec2 {
	h := hash2(p).Sub(Vec2{0.5, 0.5})
	return h.Normalize()
}

// smoothstep5 (quintic) applied component-wise.
// sub*sub*sub*(10 + sub*(-15 + 6*sub))
func quintic2(sub Vec2) Vec2 {
	sx := sub[0]
	sy := sub[1]
	qx := sx * sx * sx * (10.0 + sx*(-15.0+6.0*sx))
	qy := sy * sy * sy * (10.0 + sy*(-15.0+6.0*sy))
	return Vec2{qx, qy}
}

// mix linear interpolation.
func mix(a, b, t float64) float64 { return a*(1-t) + b*t }

// Perlin2 computes 2D Perlin noise for point p (continuous coordinates).
// Output is approximately in [0,1] after scaling (0.7) and biasing (+0.5),
// following Xor’s note (gradients range ~[-sqrt(2), +sqrt(2)]).
func Perlin2(p Vec2) float64 {
	// Cell integer coordinates and sub-cell position.
	cx, cy := math.Floor(p[0]), math.Floor(p[1])
	cell := Vec2{cx, cy}
	sub := p.Sub(cell) // fractional part in [0,1)^2

	// Offsets to cell corners (00,10,01,11).
	off00 := Vec2{0, 0}
	off10 := Vec2{1, 0}
	off01 := Vec2{0, 1}
	off11 := Vec2{1, 1}

	// Random gradient directions at corners.
	dir00 := hash2Norm(cell.Add(off00))
	dir10 := hash2Norm(cell.Add(off10))
	dir01 := hash2Norm(cell.Add(off01))
	dir11 := hash2Norm(cell.Add(off11))

	// Distance from sample to corners.
	d00 := off00.Sub(sub)
	d10 := off10.Sub(sub)
	d01 := off01.Sub(sub)
	d11 := off11.Sub(sub)

	// Gradient dot products (projection along random axes).
	g00 := dir00.Dot(d00)
	g10 := dir10.Dot(d10)
	g01 := dir01.Dot(d01)
	g11 := dir11.Dot(d11)

	// Quintic interpolation weights.
	q := quintic2(sub)

	// Interpolate horizontally, then vertically.
	hor0 := mix(g00, g10, q[0])
	hor1 := mix(g01, g11, q[0])

	// Scale to roughly [0,1]: *0.7 (≈ sqrt(0.5)) and bias +0.5
	return mix(hor0, hor1, q[1])*0.7 + 0.5
}
