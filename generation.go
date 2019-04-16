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
)

// grid of 2D gradient vectors
var gradientVectors [][]Vector

// adds perlin noise to a given section of a coordinate grid
func addPerlinNoise(grid [][]float64, minX, maxX, minY, maxY int) [][]float64 {
	xRange := maxX - minX
	yRange := maxY - minY
	vectXRange := MaxVectX - MinVectX
	vectYRange := MaxVectY - MinVectY

	rand.Seed(time.Now().UnixNano())

	// initialized the gradient vector grid
	gradientVectors = make([][]Vector, xRange+1)
	for i := range gradientVectors {
		gradientVectors[i] = make([]Vector, yRange+1)
		for j := range gradientVectors[i] {
			gradientVectors[i][j] = Vector{
				MinVectX + rand.Float64()*vectXRange,
				MinVectY + rand.Float64()*vectYRange,
			}
		}
	}

	// determine elevation a each point using the perlin() function
	for x := 0; x < xRange; x++ {
		for y := 0; y < yRange; y++ {
			grid[minX+x][minY+y] +=
				perlin(float64(x)/float64(xRange/3), float64(y)/float64(yRange/3))
		}
	}

	return grid
}

// determines a point's elevation based on
// surroundings and the gradient vector field
func perlin(x, y float64) float64 {
	x0 := int(x)
	x1 := x0 + 1
	y0 := int(y)
	y1 := y0 + 1

	sx := x - float64(x0)
	sy := y - float64(y0)

	var n0, n1, ix0, ix1 float64

	n0 = dotGridGradient(x0, y0, x, y)
	n1 = dotGridGradient(x1, y0, x, y)
	ix0 = linearInterpolate(n0, n1, sx)

	n0 = dotGridGradient(x0, y1, x, y)
	n1 = dotGridGradient(x1, y1, x, y)
	ix1 = linearInterpolate(n0, n1, sx)

	return linearInterpolate(ix0, ix1, sy)
}

// finds the dot product of two vectors
func dotGridGradient(ix, iy int, x, y float64) float64 {
	dx := x - float64(ix)
	dy := y - float64(iy)

	return dx*gradientVectors[ix][iy].x + dy*gradientVectors[ix][iy].y
}

// interpolates linearly between two points
func linearInterpolate(a0, a1, w float64) float64 {
	return (1.0-w)*a0 + w*a1
}

// adds generice random noise to a specified portion of a 2d grid
func addRandomNoise(grid [][]float64, minX, maxX, minY, maxY int) [][]float64 {
	xRange := maxX - minX
	yRange := maxY - minY

	rand.Seed(time.Now().UnixNano())

	// add or subtract a random float amount
	// from each point
	for x := 0; x < xRange; x++ {
		for y := 0; y < yRange; y++ {
			grid[x][y] += rand.Float64() * (0.005 - -0.005)
		}
	}

	return grid
}
