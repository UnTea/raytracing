package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/UnTea/raytracing/linearalgebra"
)

const width, height int = 1024, 768
var fov float64 = math.Pi / 2

func Render(spheres []linearalgebra.Sphere) {
	frameBuffer := make([]linearalgebra.Vector, width*height)
	aspectRatio := float64(width) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			normalizedX := 2.0*(float64(x)+0.5)/float64(width) - 1.0
			normalizedY := -(2.0*(float64(y)+0.5)/float64(height) - 1.0)

			filmX := normalizedX * math.Tan(fov/2.0) * aspectRatio
			filmY := normalizedY * math.Tan(fov/2.0)

			direction := linearalgebra.Normalize(linearalgebra.Vector{X: filmX, Y: filmY, Z: -1.0})
			frameBuffer[x+y*width] = linearalgebra.CastRay(spheres ,linearalgebra.Vector{X: 0.0, Y: 0.0, Z: 0.0}, direction)
		}
	}

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(255 * frameBuffer[x+y*width].X),
				G: uint8(255 * frameBuffer[x+y*width].Y),
				B: uint8(255 * frameBuffer[x+y*width].Z),
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
	ivory := linearalgebra.Material{DiffuseColor: linearalgebra.Vector{X: 0.4, Y: 0.4, Z: 0.3}}
	redRubber := linearalgebra.Material{DiffuseColor: linearalgebra.Vector{X: 0.3, Y: 0.1, Z: 0.1}}

	var spheres []linearalgebra.Sphere

	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X: -3.0, Y:  0.0, Z: -16.0}, Radius: 2.0, Material: ivory})
	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X: -1.0, Y: -1.5, Z: -12.0}, Radius: 2.0, Material: redRubber})
	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X:  1.5, Y: -0.5, Z: -18.0}, Radius: 3.0, Material: redRubber})
	spheres = append(spheres, linearalgebra.Sphere{Center: linearalgebra.Vector{X:  7.0, Y:  5.0, Z: -18.0}, Radius: 4.0, Material: ivory})

	Render(spheres)
}
