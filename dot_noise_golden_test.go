package fuzzl

import "testing"

func TestDotNoiseGolden(t *testing.T) {
	// Fixed inputs with expected outputs (generated via gen_golden.go).
	cases := []struct {
		p    Vec3
		want float64
	}{
		{Vec3{0, 0, 0}, 0.0000000000000000},
		{Vec3{1, 2, 3}, -0.6843639602299883},
		{Vec3{-1.5, 0.5, 2.0}, -0.4833420437035034},
		{Vec3{3.14159, 2.71828, 1.61803}, 0.2075428779683780},
		{Vec3{10, -10, 5}, 0.8901102473468471},
	}

	const eps = 1e-12
	for _, c := range cases {
		got := DotNoise(c.p)
		if diff := got - c.want; diff > eps || diff < -eps {
			t.Errorf("DotNoise(%v) = %v, want %v", c.p, got, c.want)
		}
	}
}
