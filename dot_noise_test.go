package fuzzl

import (
	"math/rand"
	"testing"
	"time"
)

func TestDotNoiseRange(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	const samples = 1000
	const minExpected, maxExpected = -3.0, 3.0

	for i := 0; i < samples; i++ {
		p := Vec3{
			rng.Float64()*20 - 10, // random in [-10,10]
			rng.Float64()*20 - 10,
			rng.Float64()*20 - 10,
		}
		v := DotNoise(p)

		if v < minExpected-0.001 || v > maxExpected+0.001 {
			t.Errorf("DotNoise out of expected range: got %f for p=%v", v, p)
		}

		// Log a few values for human inspection.
		if i < 5 {
			t.Logf("p=%v DotNoise(p)=%f", p, v)
		}
	}
}
