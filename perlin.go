package main

import (
	"math/rand"
	"time"
)

type Vector struct {
	x, y float64
}

var grads [][]Vector

func (w *World) addPerlinNoise(minX, maxX, minY, maxY int) {
	xRange := maxX - minX
	yRange := maxY - minY

	rand.Seed(time.Now().UnixNano())
	grads = make([][]Vector, xRange+1)
	for i := range grads {
		grads[i] = make([]Vector, yRange+1)
		for j := range grads[i] {
			grads[i][j] = Vector{-1.0 + rand.Float64()*(1.0 - -1.0),
				-1.0 + rand.Float64()*(1.0 - -1.0)}
		}
	}

	for x := 0; x < xRange; x++ {
		for y := 0; y < yRange; y++ {
			w.Elevation[minX+x][minY+y] =
				perlin(float64(x)/float64(xRange/5), float64(y)/float64(yRange/5))
		}
	}
}

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
	ix0 = lerp(n0, n1, sx)

	n0 = dotGridGradient(x0, y1, x, y)
	n1 = dotGridGradient(x1, y1, x, y)
	ix1 = lerp(n0, n1, sx)

	return lerp(ix0, ix1, sy)
}

func dotGridGradient(ix, iy int, x, y float64) float64 {

	dx := x - float64(ix)
	dy := y - float64(iy)

	return dx*grads[ix][iy].x + dy*grads[ix][iy].y
}

func lerp(a0, a1, w float64) float64 {
	return (1.0-w)*a0 + w*a1
}
