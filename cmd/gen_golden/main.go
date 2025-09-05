// gen_golden.go
//
// Helper program to generate golden values for fuzzl.DotNoise.
// Run: go run gen_golden.go
// Package main generates data for golden tests.
package main

import (
	"fmt"

	"github.com/mdhender/fuzzl"
)

func main() {
	points := []fuzzl.Vec3{
		{0, 0, 0},
		{1, 2, 3},
		{-1.5, 0.5, 2.0},
		{3.14159, 2.71828, 1.61803},
		{10, -10, 5},
	}

	for _, p := range points {
		v := fuzzl.DotNoise(p)
		fmt.Printf("{fuzzl.Vec3{%v, %v, %v}, %.16f},\n", p[0], p[1], p[2], v)
	}
}
