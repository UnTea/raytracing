package linearalgebra

import "math"

type Vector struct {
	X, Y, Z float64
}

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

type Material struct {
	DiffuseColor Vector
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

func RayIntersect(sphere Sphere, origin Vector, direction Vector, t0 *float64) bool {
	L := Subtract(sphere.Center, origin)
	tca := Dot(L, direction)
	d2 := Dot(L, L) - tca*tca

	if d2 > sphere.Radius*sphere.Radius {
		return false
	}

	thc := math.Sqrt(sphere.Radius*sphere.Radius - d2)
	*t0 = tca - thc
	t1 := tca + thc
	if *t0 < 0 {
		*t0 = t1
		return false
	}

	return true
}

func SceneIntersect(spheres []Sphere, origin Vector, direction Vector, hit *Vector, N *Vector, material *Material) bool {
	sphereDistance := math.MaxFloat64

	for i := 0; i < len(spheres); i++ {
		var distanceI float64

		if RayIntersect(spheres[i], origin, direction, &distanceI) && distanceI < sphereDistance {
			sphereDistance = distanceI
			*hit = Add(origin, Scalar(direction, distanceI))
			*N = Normalize(Subtract(*hit, spheres[i].Center))
			*material = spheres[i].Material
		}
	}

	return sphereDistance < 1000
}

func CastRay(spheres []Sphere, origin Vector, direction Vector) Vector {
	var point, N Vector
	var material Material

	if !SceneIntersect(spheres, origin, direction, &point, &N, &material) {
		return Vector{X: 0.2, Y: 0.7, Z: 0.8}
	}

	return material.DiffuseColor
}
