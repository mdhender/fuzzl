// dot_noise.go
//
// Converted from GLSL to Go.
// Original function "dot_noise" by Xor at GM Shaders.
// Source: https://mini.gmshaders.com/p/phi
//
// GM Shaders is a reader-supported publication.
// To receive new posts and support the author, consider becoming a free or paid subscriber.
// https://gmshaders.com
//
// Note: The output of DotNoise(p) typically ranges from [-3, +3].
//

package fuzzl

import "math"

// DotNoise computes the shader's dot_noise(p) for a 3D vector.
// Returns a value in roughly [-3, +3].
func DotNoise(p Vec3) float64 {
	const PHI = 1.618033988

	GOLD := [3][3]float64{
		{-0.571464913, +0.814921382, +0.096597072},
		{-0.278044873, -0.303026659, +0.911518454},
		{+0.772087367, +0.494042493, +0.399753815},
	}

	v1 := mat3MulVec3(GOLD, p)
	v2 := mat3MulVec3(mat3Transpose(GOLD), p).MulScalar(PHI)

	c := Vec3{math.Cos(v1[0]), math.Cos(v1[1]), math.Cos(v1[2])}
	s := Vec3{math.Sin(v2[0]), math.Sin(v2[1]), math.Sin(v2[2])}

	return c.Dot(s)
}

func mat3MulVec3(m [3][3]float64, v Vec3) Vec3 {
	return Vec3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2],
	}
}

func mat3Transpose(m [3][3]float64) [3][3]float64 {
	return [3][3]float64{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}
