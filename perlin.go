// Go implementation of Ken Perlin's Improved noise
package main

import (
	"math/rand"
	"time"
)

// representation of a 2D vector
type Vector struct {
	x, y float64
}

const (
	MinVectX float64 = -1.0 // minimum vector length in x direction
	MaxVectX float64 = 1.0  // maximum vector length in x direction
	MinVectY float64 = -1.0 // minimum vector length in y direction
	MaxVectY float64 = 1.0  // meximum vector length in y direction
	Scale    float64 = 5.0  // level of detail in the noise, higher is more
)

// grid of 2D gradient vectors
var gradientField [][]Vector

// adds perlin noise to a given section of a coordinate grid
func addPerlinNoise(grid [][]float64, minX, maxX, minY, maxY int) [][]float64 {
	vectXRange := MaxVectX - MinVectX
	vectYRange := MaxVectY - MinVectY

	rand.Seed(time.Now().UnixNano())

	// initialized the gradient vector grid
	gradientField = make([][]Vector, maxX-minX+1)
	for i := range gradientField {
		gradientField[i] = make([]Vector, maxY-minY+1)
		for j := range gradientField[i] {
			gradientField[i][j] = Vector{
				MinVectX + rand.Float64()*vectXRange,
				MinVectY + rand.Float64()*vectYRange,
			}
		}
	}

	// determine elevation a each point using the perlin() function
	xScale := float64(maxX-minX) / Scale
	yScale := float64(maxY-minY) / Scale
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			grid[x][y] += perlin(float64(x)/xScale, float64(y)/yScale)
		}
	}

	return grid
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
	return (1.0-w)*a0 + w*a1
}
