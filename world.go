package main

import (
	"github.com/karledramberg/ymir/procgen"
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
	world.drawMap(ElevationView)
	world.name()

	return
}

func (world *World) generate() {
	world.Elevation = procgen.AddPerlinNoise(world.Elevation, 0, world.Width, 0, world.Height)
	world.Elevation = procgen.AddRandomNoise(world.Elevation, 0, world.Width, 0, world.Height)
}

func (world *World) name() {
	world.Name = "New World"
}
