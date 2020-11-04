package main

import (
	"github.com/UnTea/raytracing/linearalgebra"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

const width, height int = 1024, 768

var fov float64 = math.Pi / 3.15

var sampleCount int = 24

func ACESFilm(v linearalgebra.Vector) linearalgebra.Vector {
	a := 2.51
	b := linearalgebra.Vector{X: 0.03, Y: 0.03, Z: 0.03}
	c := 2.43
	d := linearalgebra.Vector{X: 0.59, Y: 0.59, Z: 0.59}
	e := linearalgebra.Vector{X: 0.14, Y: 0.14, Z: 0.14}
	nominator := linearalgebra.Multiplication(v, linearalgebra.Add(linearalgebra.MulOnScalar(v, a), b))
	denominator := linearalgebra.Add(linearalgebra.Multiplication(v, linearalgebra.Add(linearalgebra.MulOnScalar(v, c), d)), e)
	return linearalgebra.Division(nominator, denominator)
}

func Render(spheres []linearalgebra.Sphere, lights []linearalgebra.Light) {
	frameBuffer := make([]linearalgebra.Vector, width*height)
	aspectRatio := float64(width) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var sum linearalgebra.Vector

			for i := 0; i < sampleCount; i++ {
				normalizedX := 2.0*(float64(x)+rand.Float64())/float64(width) - 1.0
				normalizedY := -(2.0*(float64(y)+rand.Float64())/float64(height) - 1.0)

				filmX := normalizedX * math.Tan(fov/2.0) * aspectRatio
				filmY := normalizedY * math.Tan(fov/2.0)

				direction := linearalgebra.Normalize(linearalgebra.Vector{X: filmX, Y: filmY, Z: -1.0})
				sum = linearalgebra.Add(sum, linearalgebra.CastRay(spheres, lights, linearalgebra.Vector{X: 0.0, Y: 0.0, Z: 0.0}, direction))
			}

			sum = linearalgebra.DivOnScalar(sum, float64(sampleCount))

			frameBuffer[x+y*width] = sum
		}
	}

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			filmFrameBuffer := ACESFilm(frameBuffer[x+y*width])

			img.Set(x, y, color.NRGBA{
				R: uint8(255 * math.Max(0.0, math.Min(1.0, filmFrameBuffer.X))),
				G: uint8(255 * math.Max(0.0, math.Min(1.0, filmFrameBuffer.Y))),
				B: uint8(255 * math.Max(0.0, math.Min(1.0, filmFrameBuffer.Z))),
				A: 255,
			})
		}
	}

	f, err := os.Create("raytracing.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		_ = f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Orange := linearalgebra.Material{DiffuseColor: linearalgebra.Vector{X: 1.0, Y: 0.30980, Z: 0.0}, SpecularExponent: 50}
	Green := linearalgebra.Material{DiffuseColor: linearalgebra.Vector{X: 0.45673, Y: 1.0, Z: 0.38905}, SpecularExponent: 30}
	Purple := linearalgebra.Material{DiffuseColor: linearalgebra.Vector{X: 0.30196, Y: 0.21960, Z: 0.38030}, SpecularExponent: 20}
	Turquoise := linearalgebra.Material{DiffuseColor: linearalgebra.Vector{X: 0.25098, Y: 0.87843, Z: 0.81568}, SpecularExponent: 10}

	var spheres []linearalgebra.Sphere

	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X: -3.0, Y: 0.0, Z: -16.0}, Radius: 2.0, Material: Orange})
	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X: -1.0, Y: -1.5, Z: -12.0}, Radius: 2.0, Material: Turquoise})
	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X: 1.5, Y: -0.5, Z: -18.0}, Radius: 3.0, Material: Purple})
	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X: 7.0, Y: 5.0, Z: -18.0}, Radius: 4.0, Material: Green})

	var lights []linearalgebra.Light

	lights = append(lights, linearalgebra.Light{Position: linearalgebra.Vector{X: -20, Y: 20, Z: 20}, Intensity: 1.2})
	lights = append(lights, linearalgebra.Light{Position: linearalgebra.Vector{X: 30, Y: 50, Z: -25}, Intensity: 0.1})
	lights = append(lights, linearalgebra.Light{Position: linearalgebra.Vector{X: 30, Y: 20, Z: 30}, Intensity: 0.5})

	Render(spheres, lights)
}
