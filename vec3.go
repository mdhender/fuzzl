// vec3.go
//
// Package fuzzl provides playful, shader-inspired noise generators in Go.
// Original inspiration from Xor at GM Shaders.
// Source: https://mini.gmshaders.com/p/phi
//
// GM Shaders is a reader-supported publication.
// To support the author, consider becoming a free or paid subscriber:
// https://gmshaders.com
package fuzzl

// Vec3 is a simple 3D vector type backed by [3]float64.
// It's used as the standard input for fuzzl noise functions.
type Vec3 [3]float64

// Add returns v + u.
func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v[0] + u[0], v[1] + u[1], v[2] + u[2]}
}

// Sub returns v - u.
func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v[0] - u[0], v[1] - u[1], v[2] - u[2]}
}

// MulScalar returns v * s.
func (v Vec3) MulScalar(s float64) Vec3 {
	return Vec3{v[0] * s, v[1] * s, v[2] * s}
}

// Dot returns the dot product of v · u.
func (v Vec3) Dot(u Vec3) float64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}
