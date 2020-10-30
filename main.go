package main

import (
	"github.com/UnTea/raytracing/linearmath"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const width, height int = 1024, 768
var fieldOfView float64 = math.Ceil(math.Pi/2)
var fov int = int(fieldOfView)

func Render(sphere linearmath.Sphere) {
	frameBuffer := make([]linearmath.Vector, width*height)

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			x :=  (2.0*(float64(i) + 0.5)/float64(width)  - 1.0)*math.Tan(float64(fov)/2.0)*float64(width)/float64(height)
			y := -(2.0*(float64(j) + 0.5)/float64(height) - 1.0)*math.Tan(float64(fov)/2.0)
			direction := linearmath.Normalize(linearmath.Vector{X: x, Y: y, Z: -1.0}, 1.0)
			frameBuffer[i+j*width] = linearmath.CastRay(linearmath.Vector{X: 0.0, Y: 0.0, Z: 0.0}, direction, sphere)
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

	f, err := os.Create("Picture.png")
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
	sphere := linearmath.Sphere{Center: linearmath.Vector{X: -3.0, Y: 0.0, Z: -16.0},
		Radius: 2}
	Render(sphere)
}
