package linearmath

import "math"

type Vector struct {
	X, Y, Z float64
}

type Sphere struct {
	Center Vector
	Radius float64
}

func Length(v Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func Normalize(v Vector) Vector {
	reciprocal := 1.0 / Length(v)
	return Scalar(v, reciprocal)
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

func Scalar(v Vector, number float64) Vector {
	return Vector{
		X: v.X * number,
		Y: v.Y * number,
		Z: v.Z * number,
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

func RayIntersect(sphere Sphere, origin Vector, direction Vector, t0 float64) bool {
	L := Subtract(sphere.Center, origin)
	tca := Dot(L, direction)
	d2 := Dot(L, L) - tca*tca

	if d2 > sphere.Radius*sphere.Radius {
		return false
	}

	thc := math.Sqrt(sphere.Radius*sphere.Radius - d2)
	t0 = tca - thc
	t1 := tca + thc
	if t0 < 0 {
		t0 = t1
		return false
	}

	return true
}

func CastRay(origin Vector, direction Vector, sphere Sphere) Vector {
	sphereDistance := math.MaxFloat64
	if !RayIntersect(sphere, origin, direction, sphereDistance) {
		return Vector{0.2, 0.7, 0.8}
	}

	return Vector{0.4, 0.4, 0.3}
}
