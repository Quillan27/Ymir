package main

import (
	"math/rand"
	"time"
)

func (world *World) addRandomNoise(minX, maxX, minY, maxY int) {
	xRange := maxX - minX
	yRange := maxY - minY

	rand.Seed(time.Now().UnixNano())
	for x := 0; x < xRange; x++ {
		for y := 0; y < yRange; y++ {
			world.Elevation[x][y] += rand.Float64() * (0.005 - -0.005)
		}
	}
}
