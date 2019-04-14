package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type MapView int

const (
	ElevationView MapView = iota
	ClimateView
	PoliticalView
	BiomeView
)

// creates a map image based on the world and provided mapview
func (world *World) drawMap(mapView MapView) {
	world.Map = *image.NewRGBA(image.Rect(0, 0, world.Width, world.Height))

	palette := color.Palette
	switch mapView {
	case ElevationView:
		palette = createPalette("assets/colors/elevation.png")
		for x := 0; x < world.Map.Bounds().Max.X; x++ {
			for y := 0; y < world.Map.Bounds().Max.Y; y++ {
				i := int(scale(world.Elevation[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	case ClimateView:
		palette = createPalette("assets/colors/climate.txt")
	case PoliticalView:
		palette = createPalette("assets/colors/poltical.txt")
	case BiomeView:
		palette = createPalette("assets/colors/biome.txt")
	}
}

// returns an array of colors from a .png image
func createPalette(path string) (palette color.Palette) {
	file, _ := os.Open(path)
	defer file.Close()

	image, _, _ := image.Decode(file)
	for i := 0; i < 31; i++ {
		palette = append(palette, image.At(i, 0))
	}

	return
}

// transforms a number in one range to another range
func scale(value, oldMin, oldMax, newMin, newMax float64) float64 {
	return (value-oldMin)*(newMax-newMin)/(oldMax-oldMin) + newMin
}

// saves the current worlds map to disk
func (world *World) exportMap() {
	file, err := os.Create("map.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, &world.Map)
}
