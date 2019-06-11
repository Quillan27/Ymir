package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x, y int
}

const NumberOfPolygons = 50

func (w *World) generate() {
	// Terrain

	// generate voronoi polygons

	// select seed points
	fmt.Print("Selecting cell seeds...\n")
	seeds := make([]Point, NumberOfPolygons)
	rand.Seed(time.Now().UnixNano())
	for i := range seeds {
		seeds[i].x = rand.Intn(w.Width)
		seeds[i].y = rand.Intn(w.Height)
		w.Terrain[seeds[i].x][seeds[i].y] = 2 // BLUE
	}

	// draw cell boundaries --- TESTING ONLY
	fmt.Print("Drawing cell boundaries...\n")
	for i := range seeds {
		fmt.Print("\tDrawing Cell ", i, "...\n")
		var foundAngles []float64
		for radius := 0.0; radius < 512; radius++ {
			for angle := 0.0; angle < 360.0; angle += 0.5 {
				if contains(foundAngles, angle) {
					continue
				}

				x := chomp(int(float64(radius)*math.Cos(angle))+seeds[i].x, 0, w.Width-1)
				y := chomp(int(float64(radius)*math.Sin(angle))+seeds[i].y, 0, w.Height-1)
				if closestPoint(x, y, &seeds) != i {
					foundAngles = append(foundAngles, angle)
					w.Terrain[x][y] = 1 // BLACK
				}
			}
		}
	}

	// improve points - finding centroid (average) within each polygon

	// Climate
	// Biomes
}

func closestPoint(x, y int, points *[]Point) int {
	shortestDistance := distance(x, y, (*points)[0].x, (*points)[0].y)
	candidate := 0
	for i, point := range *points {
		if distance(x, y, point.x, point.y) < shortestDistance {
			shortestDistance = distance(x, y, point.x, point.y)
			candidate = i
		}
	}
	return candidate
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func chomp(num, min, max int) int {
	if num < min {
		return min
	} else if num > max {
		return max
	}
	return num
}
