package linmath

import "math"

type Vector struct {
	X, Y, Z float64
}

func Length(v Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func Normalize(v Vector) Vector {
	reciprocal := 1.0 / Length(v)
	return MulOnScalar(v, reciprocal)
}

func Multiplication(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
	}
}

func Dot(v1 Vector, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func MulOnScalar(v Vector, number float64) Vector {
	return Vector{
		X: v.X * number,
		Y: v.Y * number,
		Z: v.Z * number,
	}
}

func DivOnScalar(v Vector, number float64) Vector {
	return Vector{
		X: v.X / number,
		Y: v.Y / number,
		Z: v.Z / number,
	}
}

func Add(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func Subtract(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func Division(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X / v2.X,
		Y: v1.Y / v2.Y,
		Z: v1.Z / v2.Z,
	}
}
