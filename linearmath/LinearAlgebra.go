package linearmath

import "math"

type Vector struct {
	X, Y, Z float64
}

type Sphere struct {
	Center Vector
	Radius float64
}

func Length(vector Vector) float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)
}

func Normalize(vector Vector) Vector {
	reciprocal := 1.0 / Length(vector)
	return MulScalar(vector, reciprocal)
}

func MulVector(vector1 Vector, vector2 Vector) Vector {
	return Vector{
		X: vector1.X * vector2.X,
		Y: vector1.Y * vector2.Y,
		Z: vector1.Z * vector2.Z,
	}
}

func Dot(vector1 Vector, vector2 Vector) float64 {
	return vector1.X*vector2.X + vector1.Y*vector2.Y + vector1.Z*vector2.Z
}

func MulScalar(vector Vector, number float64) Vector {
	return Vector{
		X: vector.X * number,
		Y: vector.Y * number,
		Z: vector.Z * number,
	}
}

func AddVector(vector1 Vector, vector2 Vector) Vector {
	return Vector{
		X: vector1.X + vector2.X,
		Y: vector1.Y + vector2.Y,
		Z: vector1.Z + vector2.Z,
	}
}

func SubVector(vector1 Vector, vector2 Vector) Vector {
	return Vector{
		X: vector1.X - vector2.X,
		Y: vector1.Y - vector2.Y,
		Z: vector1.Z - vector2.Z,
	}
}

func RayIntersect(sphere Sphere, origin Vector, direction Vector, t0 float64) bool {
	L := SubVector(sphere.Center, origin)
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
