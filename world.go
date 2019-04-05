package main

import (
	"image"
)

type World struct {
	Grid [][][]float64
	Map  image.RGBA
	Name string
}

func newWorld(width, height int) (w *World) {
	w = new(World)
	w.Grid = make([][][]float64, width)
	for x := range w.Grid {
		w.Grid[x] = make([][]float64, height)
		for y := range w.Grid[x] {
			w.Grid[x][y] = make([]float64, 4)
		}
	}
	w.generate()
	w.drawMap(ELEVATION)
	w.name()

	return
}

func (w *World) generate() {
	w.addPerlinNoise(0, 0, len(w.Grid), len(w.Grid[0]))
}

func (w *World) name() {
	w.Name = "New World"
}
