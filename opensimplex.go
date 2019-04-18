// Go implementation of Kurt Spencer's OpenSimplex noise
package main

import "math/rand"

const (
	StretchConstant = -0.211324865405187 // (1/Math.sqrt(2+1)-1)/2
	SquishConstant  = 0.366025403784439  // (Math.sqrt(2+1)-1)/2
	NormConstant    = 47
	DefaultSeed     = 0
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
func addOpenSimplexNoise(grid [][]float64, minX, maxX, minY, maxY int) [][]float64 {
	perm = make([]int16, 256)

	source := make([]int16, 256)
	for i := int16(0); i < 256; i++ {
		source[i] = i
	}

	seed := rand.Int63()
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407

	for i := 255; i >= 0; i-- {
		seed = seed*6364136223846793005 + 1442695040888963407
		r := int((seed + 31) % int64(i+1))
		if r < 0 {
			r += (i + 1)
		}
		perm[i] = source[r]
		source[r] = source[i]
	}

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			grid[x][y] = opensimplex(float64(x), float64(y))
		}
	}
	return grid
}

func opensimplex(x, y float64) float64 {
	stretchOffset := (x + y) * StretchConstant
	xs := x + stretchOffset
	ys := y + stretchOffset

	// floor to get grid coords of the rhombus
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

	var value float64

	// 0,1
	dx1 := dx0 - 1 - SquishConstant
	dy1 := dy0 - 0 - SquishConstant
	attn1 := 2 - dx1*dx1 - dy1*dy1
	if attn1 > 0 {
		attn1 *= attn1
		value += attn1 * attn1 * extrapolate(xsb+1, ysb+0, dx1, dy1)
	}

	// 1,0
	dx2 := dx0 - 0 - SquishConstant
	dy2 := dy0 - 1 - SquishConstant
	attn2 := 2 - dx2*dx2 - dy2*dy2
	if attn2 > 0 {
		attn2 *= attn2
		value += attn2 * attn2 * extrapolate(xsb+1, ysb+0, dx2, dy2)
	}

	var xsv_ext, ysv_ext int
	var dx_ext, dy_ext float64

	if inSum <= 1 {
		zins := 1 - inSum
		if zins > xins || zins > yins {
			if xins > yins {
				xsv_ext = xsb + 1
				ysv_ext = ysb - 1
				dx_ext = dx0 - 1
				dy_ext = dy0 + 1
			} else {
				xsv_ext = xsb - 1
				ysv_ext = ysb + 1
				dx_ext = dx0 + 1
				dy_ext = dy0 - 1
			}
		} else {
			xsv_ext = xsb + 1
			ysv_ext = ysb + 1
			dx_ext = dx0 - 1 - 2*SquishConstant
			dy_ext = dy0 - 1 - 2*SquishConstant
		}
	} else {
		zins := 2 - inSum
		if zins < xins || zins < yins {
			if xins > yins {
				xsv_ext = xsb + 2
				ysv_ext = ysb + 0
				dx_ext = dx0 - 2 - 2*SquishConstant
				dy_ext = dy0 + 0 - 2*SquishConstant
			} else {
				xsv_ext = xsb + 0
				ysv_ext = ysb + 2
				dx_ext = dx0 + 0 - 2*SquishConstant
				dy_ext = dy0 - 2 - 2*SquishConstant
			}
		} else {
			dx_ext = dx0
			dy_ext = dy0
			xsv_ext = xsb
			ysv_ext = ysb
		}
		xsb += 1
		ysb += 1
		dx0 = dx0 - 1 - 2*SquishConstant
		dy0 = dy0 - 1 - 2*SquishConstant
	}

	attn0 := 2 - dx0*dx0 - dy0*dy0
	if attn0 > 0 {
		attn0 *= attn0
		value += attn0 * attn0 * extrapolate(xsb, ysb, dx0, dy0)
	}

	attn_ext := 2 - dx_ext*dx_ext - dy_ext*dy_ext
	if attn_ext > 0 {
		attn_ext *= attn_ext
		value += attn_ext * attn_ext * extrapolate(xsv_ext, ysv_ext, dx_ext, dy_ext)
	}

	return value / NormConstant
}

func extrapolate(xsb, ysb int, dx, dy float64) float64 {
	index := perm[(int(perm[xsb&0xFF])+ysb)&0xFF] & 0x0E
	return float64(gradients[index])*dx + float64(gradients[index+1])*dy
}
