// Go implementation of Ken Perlin's Improved noise
package main

import (
	"math/rand"
	"time"
)

// Vector represents a 2D vector
type Vector struct {
	x, y float64
}

const (
	// MinVectX is the minimum vector length in x direction
	MinVectX float64 = -1.0
	// MaxVectX is the maximum vector length in x direction
	MaxVectX float64 = 1.0
	// MinVectY is the minimum vector length in y direction
	MinVectY float64 = -1.0
	// MaxVectY is the maximum vector length in y direction
	MaxVectY float64 = 1.0
)

// gradientField is a grid of 2D vectors
var gradientField [][]Vector

// addPerlinNoise adds perlin noise to a given section of a coordinate grid
func addPerlinNoise(grid *[][]float64, octaves int, persistence float64) {
	gridWidth := len(*grid) + 1
	gridHeight := len((*grid)[0]) + 1

	// initialized the gradient vector grid
	rand.Seed(time.Now().UnixNano())
	gradientField = make([][]Vector, gridWidth)
	for x := range gradientField {
		gradientField[x] = make([]Vector, gridHeight)
		for y := range gradientField[x] {
			gradientField[x][y] = Vector{
				MinVectX + rand.Float64()*(MaxVectX-MinVectX),
				MinVectY + rand.Float64()*(MaxVectY-MinVectY),
			}
		}
	}

	// determine elevation at each point using the perlin() function
	for x := range *grid {
		for y := range (*grid)[0] {
			value := 0.0
			frequency := 1.0
			amplitude := 8.0
			maxValue := -1.0
			for i := 0; i < octaves; i++ {
				value += perlin(float64(x)/frequency,
					float64(y)/frequency) * amplitude
				maxValue += amplitude
				amplitude *= persistence
				frequency *= 2
			}
			(*grid)[x][y] += value / maxValue
		}
	}
}

// perlin determines a point's elevation based on
// surroundings and the gradient vector field
// BUG, noise is very grid like - might just be perlin
func perlin(x, y float64) (value float64) {
	// determine the gradient vector to interpolate with
	// by snapping to gradientField coordinates
	x0 := int(x) // left side
	x1 := x0 + 1 // right side
	y0 := int(y) // top side
	y1 := y0 + 1 // bottom side

	sx := x - float64(x0) // determine x-axis interpolation weight
	sy := y - float64(y0) // determine y-axis interpolation weight

	n0 := dotWithGradient(x0, y0, x, y) // top left corner
	n1 := dotWithGradient(x1, y0, x, y) // top right corner
	ix0 := interpolate(n0, n1, sx)      // interpolate along the top

	n0 = dotWithGradient(x0, y1, x, y) // bottom left corner
	n1 = dotWithGradient(x1, y1, x, y) // bottom right corner
	ix1 := interpolate(n0, n1, sx)     // interpolate along the bottom

	value = interpolate(ix0, ix1, sy) // interpolate along the y-axis

	return
}

// dotWithGradient finds the dot product of the
// gradient vector and the distance vector
func dotWithGradient(ix, iy int, x, y float64) float64 {
	dx := x - float64(ix)
	dy := y - float64(iy)

	return dx*gradientField[ix][iy].x + dy*gradientField[ix][iy].y
}

// interpolate interpolates smoothly between two points
func interpolate(a0, a1, w float64) float64 {
	return a0 + (w*w*w*(w*(w*(w*6-15)+10)))*(a1-a0)
}
