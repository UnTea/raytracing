package main

import (
	"github.com/UnTea/raytracing/linmath"
	"github.com/UnTea/raytracing/raytracing"
)

type Vector = linmath.Vector
type Sphere = raytracing.Sphere
type Light = raytracing.Light
type Material = raytracing.Material

func main() {
	Orange := Material{DiffuseColor: Vector{X: 1.0, Y: 0.30980, Z: 0.0}, SpecularExponent: 50}
	Green := Material{DiffuseColor: Vector{X: 0.45673, Y: 1.0, Z: 0.38905}, SpecularExponent: 30}
	Purple := Material{DiffuseColor: Vector{X: 0.30196, Y: 0.21960, Z: 0.38030}, SpecularExponent: 20}
	Turquoise := Material{DiffuseColor: Vector{X: 0.25098, Y: 0.87843, Z: 0.81568}, SpecularExponent: 10}

	var spheres []Sphere

	spheres = append(spheres, Sphere{Center: Vector{X: -3.0, Y: 0.0, Z: -16.0}, Radius: 2.0, Material: Orange})
	spheres = append(spheres, Sphere{Center: Vector{X: -1.0, Y: -1.5, Z: -12.0}, Radius: 2.0, Material: Turquoise})
	spheres = append(spheres, Sphere{Center: Vector{X: 1.5, Y: -0.5, Z: -18.0}, Radius: 3.0, Material: Purple})
	spheres = append(spheres, Sphere{Center: Vector{X: 7.0, Y: 5.0, Z: -18.0}, Radius: 4.0, Material: Green})

	var lights []Light

	lights = append(lights, Light{Position: Vector{X: -20, Y: 20, Z: 20}, Intensity: 1.2})
	lights = append(lights, Light{Position: Vector{X: 30, Y: 50, Z: -25}, Intensity: 0.9})
	lights = append(lights, Light{Position: Vector{X: 30, Y: 20, Z: 30}, Intensity: 0.5})

	raytracing.Render(spheres, lights)
}
