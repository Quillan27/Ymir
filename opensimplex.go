// Go implementation of Kurt Spencer's OpenSimplex noise
package main

import "math/rand"

const (
	// StretchConstant is used to strech a orthoganal grid
	// into a diamond shaped one, later split into triangles
	// (1/Math.sqrt(2+1)-1)/2
	StretchConstant = -0.211324865405187

	// SquishConstant is anothor constant used to make
	// a diamond grid from an orthoganal one
	// (Math.sqrt(2+1)-1)/2
	SquishConstant = 0.366025403784439

	// NormConstant is used to make the noise values more even
	// and usable
	NormConstant = 47
)

// gradients approximates the direction to
// the vertices of an octagon from the center
var gradients = []int8{
	5, 2, 2, 5,
	-5, 2, -2, 5,
	5, -2, 2, -5,
	-5, -2, -2, -5,
}

var perm []int16

// addOpenSimplexNoise will add Open Simplex Noise
// to the specified range of the provided 2D grid
// (BUG) checker pattern on high x or high y values without a corresponding high x or y
func addOpenSimplexNoise(grid [][]float64, minX, maxX, minY, maxY int) [][]float64 {
	perm = make([]int16, 256)

	source := make([]int16, 256)
	for i := range source {
		source[i] = int16(i)
	}

	seed := rand.Int63()
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407

	for i := 255; i >= 0; i-- {
		seed = seed*6364136223846793005 + 1442695040888963407
		r := int((seed + 31) % int64(i+1))
		if r < 0 {
			r += i + 1
		}
		perm[i] = source[r]
		source[r] = source[i]
	}

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			grid[x][y] = opensimplex(float64(x)/10, float64(y)/10)
		}
	}
	return grid
}

func opensimplex(x, y float64) float64 {
	stretchOffset := (x + y) * StretchConstant
	xs := float64(x + stretchOffset)
	ys := float64(y + stretchOffset)

	xsb := int(xs)
	ysb := int(ys)

	squishOffset := float64(xsb+ysb) * SquishConstant
	xb := float64(xsb) + squishOffset
	yb := float64(ysb) + squishOffset

	xins := xs - float64(xsb)
	yins := ys - float64(ysb)

	inSum := xins + yins

	dx0 := x - xb
	dy0 := y - yb

	var dxExt, dyExt float64
	var xsvExt, ysvExt int

	var value float64

	dx1 := dx0 - 1 - SquishConstant
	dy1 := dy0 - 0 - SquishConstant
	attn1 := 2 - dx1*dx1 - dy1*dy1
	if attn1 > 0 {
		attn1 *= attn1
		value += attn1 * attn1 * extrapolate(xsb+1, ysb+0, dx1, dy1)
	}

	dx2 := dx0 - 0 - SquishConstant
	dy2 := dy0 - 1 - SquishConstant
	attn2 := 2 - dx2*dx2 - dy2*dy2
	if attn2 > 0 {
		attn2 *= attn2
		value += attn2 * attn2 * extrapolate(xsb+0, ysb+1, dx2, dy2)
	}

	if inSum <= 1 {
		zins := 1 - inSum
		if zins > xins || zins > yins {
			if xins > yins {
				xsvExt = xsb + 1
				ysvExt = ysb - 1
				dxExt = dx0 - 1
				dyExt = dy0 + 1
			} else {
				xsvExt = xsb - 1
				ysvExt = ysb + 1
				dxExt = dx0 + 1
				dyExt = dy0 - 1
			}
		} else {
			xsvExt = xsb + 1
			ysvExt = ysb + 1
			dxExt = dx0 - 1 - 2*SquishConstant
			dyExt = dy0 - 1 - 2*SquishConstant
		}
	} else {
		zins := 2 - inSum
		if zins < xins || zins < yins {
			if xins > yins {
				xsvExt = xsb + 2
				ysvExt = ysb + 0
				dxExt = dx0 - 2 - 2*SquishConstant
				dyExt = dy0 + 0 - 2*SquishConstant
			} else {
				xsvExt = xsb + 0
				ysvExt = ysb + 2
				dxExt = dx0 + 0 - 2*SquishConstant
				dyExt = dy0 - 2 - 2*SquishConstant
			}
		} else {
			dxExt = dx0
			dyExt = dy0
			xsvExt = xsb
			ysvExt = ysb
		}
		xsb++
		ysb++
		dx0 = dx0 - 1 - 2*SquishConstant
		dy0 = dy0 - 1 - 2*SquishConstant
	}

	attn0 := 2 - dx0*dx0 - dy0*dy0
	if attn0 > 0 {
		attn0 *= attn0
		value += attn0 * attn0 * extrapolate(xsb, ysb, dx0, dy0)
	}

	attnExt := 2 - dxExt*dxExt - dyExt*dyExt
	if attnExt > 0 {
		attnExt *= attnExt
		value += attnExt * attnExt * extrapolate(xsvExt, ysvExt, dxExt, dyExt)
	}

	return value / NormConstant
}

func extrapolate(xsb, ysb int, dx, dy float64) float64 {
	index := perm[(int(perm[xsb&0xFF])+ysb)&0xFF] & 0x0E
	return float64(gradients[index])*dx + float64(gradients[index+1])*dy
}
