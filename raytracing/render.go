package raytracing

import (
	"github.com/UnTea/raytracing/linmath"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

const width, height, sampleCount int = 1024, 768, 2048
const fov float64 = math.Pi / 3.15

func Render(spheres []Sphere, lights []Light) {
	frameBuffer := make([]Vector, width*height)
	aspectRatio := float64(width) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var sum Vector

			for i := 0; i < sampleCount; i++ {
				normalizedX := 2.0*(float64(x)+rand.Float64())/float64(width) - 1.0
				normalizedY := -(2.0*(float64(y)+rand.Float64())/float64(height) - 1.0)

				filmX := normalizedX * math.Tan(fov/2.0) * aspectRatio
				filmY := normalizedY * math.Tan(fov/2.0)

				direction := linmath.Normalize(Vector{X: filmX, Y: filmY, Z: -1.0})
				sum = linmath.Add(sum, CastRay(spheres, lights, Vector{X: 0.0, Y: 0.0, Z: 0.0}, direction))
			}

			sum = linmath.DivOnScalar(sum, float64(sampleCount))
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

	f, err := os.Create("render.png")
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
