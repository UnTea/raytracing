package raytracing

import (
	"github.com/UnTea/raytracing/linmath"
	"math"
)

type Vector = linmath.Vector

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

type Material struct {
	DiffuseColor     Vector
	SpecularExponent float64
}

type Light struct {
	Position  Vector
	Intensity float64
}

func ACESFilm(v Vector) Vector {
	a := 2.51
	b := linmath.Vector{X: 0.03, Y: 0.03, Z: 0.03}
	c := 2.43
	d := linmath.Vector{X: 0.59, Y: 0.59, Z: 0.59}
	e := linmath.Vector{X: 0.14, Y: 0.14, Z: 0.14}
	nominator := linmath.Multiplication(v, linmath.Add(linmath.MulOnScalar(v, a), b))
	denominator := linmath.Add(linmath.Multiplication(v, linmath.Add(linmath.MulOnScalar(v, c), d)), e)
	return linmath.Division(nominator, denominator)
}

func Reflect(I Vector, N Vector) Vector {
	return linmath.Subtract(I, linmath.MulOnScalar(linmath.MulOnScalar(N, 2.0), linmath.Dot(I, N)))
}

func CastRay(spheres []Sphere, lights []Light, origin Vector, direction Vector) Vector {
	var point, N Vector
	var material Material

	if !SceneIntersect(spheres, origin, direction, &point, &N, &material) {
		return Vector{X: 0.2, Y: 0.7, Z: 0.8}
	}

	var diffuseLightIntensity, specularLightIntensity float64

	for i := 0; i < len(lights); i++ {
		lightDirection := linmath.Normalize(linmath.Subtract(lights[i].Position, point))
		lightDistance := linmath.Length(linmath.Subtract(lights[i].Position, point))
		var shadowOrigin Vector

		if linmath.Dot(lightDirection, N) < 0 {
			shadowOrigin = linmath.Subtract(point, linmath.MulOnScalar(N, 1e-3))
		} else {
			shadowOrigin = linmath.Add(point, linmath.MulOnScalar(N, 1e-3))
		}

		var shadowPoint, shadowN Vector
		var tempMaterial Material

		if SceneIntersect(spheres, shadowOrigin, lightDirection, &shadowPoint, &shadowN, &tempMaterial) && (linmath.Length(linmath.Subtract(shadowPoint, shadowOrigin)) < lightDistance) {
			continue
		}

		diffuseLightIntensity += lights[i].Intensity * math.Max(0.0, linmath.Dot(lightDirection, N))
		specularLightIntensity += math.Pow(math.Max(0.0, linmath.Dot(Reflect(lightDirection, N), direction)), material.SpecularExponent) * lights[i].Intensity
	}

	first := linmath.MulOnScalar(material.DiffuseColor, diffuseLightIntensity)
	second := linmath.MulOnScalar(Vector{X: 1.0, Y: 1.0, Z: 1.0}, specularLightIntensity)

	return linmath.Add(first, second)
}

func SceneIntersect(spheres []Sphere, origin Vector, direction Vector, hit *Vector, N *Vector, material *Material) bool {
	sphereDistance := math.MaxFloat64

	for i := 0; i < len(spheres); i++ {
		var distanceToI float64

		if RayIntersect(spheres[i], origin, direction, &distanceToI) && distanceToI < sphereDistance {
			sphereDistance = distanceToI
			*hit = linmath.Add(origin, linmath.MulOnScalar(direction, distanceToI))
			*N = linmath.Normalize(linmath.Subtract(*hit, spheres[i].Center))
			*material = spheres[i].Material
		}
	}

	return sphereDistance < 1000
}

func RayIntersect(sphere Sphere, origin Vector, direction Vector, t0 *float64) bool {
	L := linmath.Subtract(sphere.Center, origin)
	tca := linmath.Dot(L, direction)
	d2 := linmath.Dot(L, L) - tca*tca

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
