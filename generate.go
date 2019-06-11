package main

import (
	"math"
	"math/rand"
)

type Point struct {
	x, y int
}

const NumberOfPolygons = 50

func (w *World) generate() {
	// Terrain

	// generate voronoi polygons

	// select seed points
	var seeds [NumberOfPolygons]Point
	for _, seed := range seeds {
		seed.x = rand.Intn(w.Width)
		seed.y = rand.Intn(w.Height)
		w.Terrain[seed.x][seed.y] = 2 // BLUE
	}

	// improve points - finding centroid (average) within each polygon

	// Climate
	// Biomes
}

func findClosesetPoint(x, y int, points *[]Point) Point {
	shortestDistance := distance(x, y, (*points)[0].x, (*points)[0].y)
	closestPoint := (*points)[0]
	for i, point := range *points {
		if distance(x, y, point.x, point.y) < shortestDistance {
			shortestDistance = distance(x, y, point.x, point.y)
			closestPoint = (*points)[i]
		}
	}
	return closestPoint
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}
