// vec2.go
//
// Package fuzzl — shader-inspired noise generators in Go.
// Attribution: Xor / GM Shaders  — https://mini.gmshaders.com/p/phi
// Support GM Shaders: https://gmshaders.com

package fuzzl

import "math"

// Vec2 is a minimal 2D vector backed by [2]float64.
type Vec2 [2]float64

func (v Vec2) Add(u Vec2) Vec2          { return Vec2{v[0] + u[0], v[1] + u[1]} }
func (v Vec2) Sub(u Vec2) Vec2          { return Vec2{v[0] - u[0], v[1] - u[1]} }
func (v Vec2) MulScalar(s float64) Vec2 { return Vec2{v[0] * s, v[1] * s} }
func (v Vec2) Dot(u Vec2) float64       { return v[0]*u[0] + v[1]*u[1] }
func (v Vec2) Length() float64          { return math.Hypot(v[0], v[1]) }
func (v Vec2) Normalize() Vec2 {
	if n := v.Length(); n > 0 {
		return Vec2{v[0] / n, v[1] / n}
	}
	return v
}

func fract(x float64) float64 { return x - math.Floor(x) }
