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

// grid of 2D gradient vectors
var gradientField [][]Vector

// adds perlin noise to a given section of a coordinate grid
func addPerlinNoise(grid *[][]float64, octaves int, persistence float64) {
	gridWidth := len(*grid)
	gridHeight := len((*grid)[0])
	vectXRange := MaxVectX - MinVectX
	vectYRange := MaxVectY - MinVectY

	rand.Seed(time.Now().UnixNano())

	// initialized the gradient vector grid
	gradientField = make([][]Vector, gridWidth+1)
	for i := range gradientField {
		gradientField[i] = make([]Vector, gridHeight+1)
		for j := range gradientField[i] {
			gradientField[i][j] = Vector{
				MinVectX + rand.Float64()*vectXRange,
				MinVectY + rand.Float64()*vectYRange,
			}
		}
	}

	// determine elevation a each point using the perlin() function
	for x := range *grid {
		for y := range (*grid)[0] {
			value := 0.0
			frequency := 1.0
			amplitude := 8.0
			maxValue := -1.0
			for i := 0; i < octaves; i++ {
				value += perlin(float64(x)/frequency, float64(y)/frequency) * amplitude
				maxValue += amplitude
				amplitude *= persistence
				frequency *= 2
			}
			(*grid)[x][y] = value / maxValue
		}
	}

}

// perlin determines a point's elevation based on
// surroundings and the gradient vector field
// BUG, noise is very grid like - might just be perlin
func perlin(x, y float64) float64 {
	// determine gradient vectors to interpolate with
	x0 := int(x)
	x1 := x0 + 1
	y0 := int(y)
	y1 := y0 + 1

	// determine interpolation weights
	sx := x - float64(x0)
	sy := y - float64(y0)

	// interpolate between grid point gradients
	var n0, n1, ix0, ix1 float64

	n0 = dotGridGradient(x0, y0, x, y)
	n1 = dotGridGradient(x1, y0, x, y)
	ix0 = linearInterpolate(n0, n1, sx)

	n0 = dotGridGradient(x0, y1, x, y)
	n1 = dotGridGradient(x1, y1, x, y)
	ix1 = linearInterpolate(n0, n1, sx)

	return linearInterpolate(ix0, ix1, sy)
}

// dotGridGradient finds the dot product of the
// gradient vector and the distance vector
func dotGridGradient(ix, iy int, x, y float64) float64 {
	dx := x - float64(ix)
	dy := y - float64(iy)

	return dx*gradientField[ix][iy].x + dy*gradientField[ix][iy].y
}

// interpolates linearly between two points
func linearInterpolate(a0, a1, w float64) float64 {
	return a0 + (w*w*w*(w*(w*(w*6-15)+10)))*(a1-a0)
}
