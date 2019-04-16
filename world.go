package main

import (
	"image"
)

type World struct {
	Elevation [][]float64 // generated 2d grid for the world's elevation
	Map       image.RGBA  // generated map from elevation and a mapview
	Name      string      // generated random name for the world
	Width     int         // world width
	Height    int         // world height
}

// generates a completely new world from scratch
func newWorld(width, height int) (world *World) {
	// create a new, blank world
	world = new(World)

	// give the world dimensions
	world.Width = width
	world.Height = height

	// initialize the elevation height map with 0s
	world.Elevation = make([][]float64, world.Width)
	for x := range world.Elevation {
		world.Elevation[x] = make([]float64, world.Height)
	}

	world.generateTerrain()
	world.drawMap(ElevationView) // elevation is the default mapview
	world.name()

	return
}

// generates a height map from scratch for a new world
func (world *World) generateTerrain() {
	world.Elevation = addPerlinNoise(world.Elevation, 0, world.Width, 0, world.Height)
	world.Elevation = addRandomNoise(world.Elevation, 0, world.Width, 0, world.Height)
}

// (TODO) generates a new random name for the world
func (world *World) name() {
	world.Name = "New World"
}
