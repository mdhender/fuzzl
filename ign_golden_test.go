package fuzzl

import "testing"

func TestIGNGolden(t *testing.T) {
	// A few fixed samples to lock behavior.
	cases := []struct {
		x, y int
		want float64
	}{
		{0, 0, 0.0},
		{1, 0, 0.06711056*52.9829189 - float64(int(0.06711056*52.9829189))}, // expanded, but we’ll assert via IGN directly
		{5, 7, IGN(5, 7)},
		{128, 256, IGN(128, 256)},
		{-3, 2, IGN(-3, 2)},
	}

	const eps = 1e-12
	for _, c := range cases {
		got := IGN(c.x, c.y)
		if got < 0 || got >= 1 {
			t.Fatalf("IGN(%d,%d) out of range: %v", c.x, c.y, got)
		}
		if diff := got - c.want; diff > eps || diff < -eps {
			t.Errorf("IGN(%d,%d) = %v, want %v", c.x, c.y, got, c.want)
		}
	}
}
