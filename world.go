package main

import (
	"image"
)

type World struct {
	Elevation [][]float64
	Map       image.RGBA
	Name      string
	Width     int
	Height    int
}

func newWorld(width, height int) (world *World) {
	world = new(World)

	world.Width = width
	world.Height = height

	world.Elevation = make([][]float64, world.Width)
	for x := range world.Elevation {
		world.Elevation[x] = make([]float64, world.Height)
	}
	world.generate()
	world.drawMap(ELEVATION)
	world.name()

	return
}

func (world *World) generate() {
	world.addPerlinNoise(0, len(world.Elevation), 0, len(world.Elevation[0]))
}

func (world *World) name() {
	world.Name = "New World"
}
