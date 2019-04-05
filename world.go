package main

import (
	"image"
)

type World struct {
	Elevation [][]float64
	Map       image.RGBA
	Name      string
}

func newWorld(width, height int) (w *World) {
	w = new(World)
	w.Elevation = make([][]float64, width)
	for x := range w.Elevation {
		w.Elevation[x] = make([]float64, height)
	}
	w.generate()
	w.drawMap(ELEVATION)
	w.name()

	return
}

func (w *World) generate() {
	w.addPerlinNoise(0, len(w.Elevation), 0, len(w.Elevation[0]))
}

func (w *World) name() {
	w.Name = "New World"
}
