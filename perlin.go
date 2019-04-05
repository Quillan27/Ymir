package main

import (
	"math/rand"
	"time"
)

func (w *World) addPerlinNoise(minX, minY, maxX, maxY int) {
	xRange := float64(maxX - minX)
	yRange := float64(maxY - minY)

	rand.Seed(time.Now().UnixNano())
	grads := make([][][]float64, int(xRange))
	for x := range grads {
		grads[x] = make([][]float64, int(yRange))
		for y := range grads[x] {
			grads[x][y] = make([]float64, 2)
			for d := range grads[x][y] {
				grads[x][y][d] = -1.0 + rand.Float64()*(1.0 - -1.0)
			}
		}
	}

	for x := 0; x < int(xRange); x++ {
		for y := 0; y < int(yRange); y++ {
			w.Grid[minX+x][minY+y][ELEVATION] = perlin(float64(x)/(xRange/(xRange*0.01)),
				float64(y)/(yRange/(yRange*0.01)), &grads)
		}
	}
}

func clamp(num, min, max float64) float64 {
	if num < min {
		return min
	} else if num > max {
		return max
	} else {
		return num
	}
}

func perlin(x, y float64, grads *[][][]float64) float64 {
	x0 := int(x)
	x1 := x0 + 1
	y0 := int(y)
	y1 := y0 + 1

	sx := x - float64(x0)
	sy := y - float64(y0)

	var n0, n1, ix0, ix1 float64

	n0 = dotGridGradient(x0, y0, x, y, grads)
	n1 = dotGridGradient(x1, y0, x, y, grads)
	ix0 = lerp(n0, n1, sx)

	n0 = dotGridGradient(x0, y1, x, y, grads)
	n1 = dotGridGradient(x1, y1, x, y, grads)
	ix1 = lerp(n0, n1, sx)

	return lerp(ix0, ix1, sy)
}

func dotGridGradient(ix, iy int, x, y float64, grads *[][][]float64) float64 {

	dx := x - float64(ix)
	dy := y - float64(iy)

	//fmt.Print("ix: ", ix, "iy: ", iy, "\n")
	return (dx*(*grads)[ix][iy][0] + dy*(*grads)[ix][iy][1])
}

func lerp(a0, a1, w float64) float64 {
	return (1.0-w)*a0 + w*a1
}
