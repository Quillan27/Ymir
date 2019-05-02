package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type World struct {
	Name    string
	Terrain [][]float64
	Width   int
	Height  int
	Map     image.RGBA
}

// MapView determines aspects like color palettes
// and how the map is generated from the terrain
type MapView uint8

const (
	// map view enumeration
	ElevationView MapView = iota
	ClimateView
	PoliticalView
	BiomeView
	TopographicView

	// asset file paths
	AssetsDir              string = "assets/"
	PaletteDir             string = AssetsDir + "palettes/"
	ElevationPalettePath   string = PaletteDir + "elevation.png"
	ClimatePalettePath     string = PaletteDir + "climate.png"
	PoliticalPalettePath   string = PaletteDir + "political.png"
	BiomePalettePath       string = PaletteDir + "biome.png"
	TopographicPalettePath string = PaletteDir + "topographic.png"
)

func newWorld(width, height int) (world *World) {
	world = new(World)

	world.Width = width
	world.Height = height

	world.Terrain = make([][]float64, world.Width)
	for x := range world.Terrain {
		world.Terrain[x] = make([]float64, world.Height)
	}

	world.generateTerrain()
	world.drawMap(TopographicView)
	world.name()
	world.exportMap()
	return
}

func (world *World) generateTerrain() {
	addPerlinNoise(&world.Terrain, 8, 2.0)
}

func (world *World) drawMap(mapView MapView) {
	world.Map = *image.NewRGBA(image.Rect(0, 0, world.Width, world.Height)) // new blank RGBA image

	palette := color.Palette{}
	switch mapView {
	case ElevationView:
		palette = createPalette(ElevationPalettePath)
	case ClimateView:
		palette = createPalette(ClimatePalettePath)
	case PoliticalView:
		palette = createPalette(PoliticalPalettePath)
	case BiomeView:
		palette = createPalette(BiomePalettePath)
	case TopographicView:
		palette = createPalette(TopographicPalettePath)
	}

	switch mapView {
	case ElevationView:
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				// ensure the elevation is between -1.0 and 1.0
				world.Terrain[x][y] = chompFloat(world.Terrain[x][y], -1.0, 1.0)

				// map the grid value to a color in the palette between 0 and 31
				color := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))

				world.Map.Set(x, y, palette[color])
			}
		}
	case ClimateView: // (TODO) algorithm for interpreting elevation for climate
		for x := 0; x < world.Map.Bounds().Max.X; x++ {
			for y := 0; y < world.Map.Bounds().Max.Y; y++ {
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	case PoliticalView: // (TODO) algorithm for interpreting political boundaries based on terrain
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	case BiomeView: // (TODO) alogrithm for interpreting biome based on climate, elevation, proximity to ocean, etc.
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	case TopographicView: // black if next to a lower elevation, white if below sealevel or otherwise
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				world.Terrain[x][y] = chompFloat(world.Terrain[x][y], -1.0, 1.0)

				color := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))

				black := 0
				white := 31
				water := 27
				wave := 17
				sealevel := 16
				if color < sealevel {
					color = water
				} else if int(scale(world.Terrain[chompInt(x+1, 0, world.Width-1)][y], -1.0, 1.0, 0.0, 31.0)) < color ||
					int(scale(world.Terrain[x][chompInt(y+1, 0, world.Height-1)], -1.0, 1.0, 0.0, 31.0)) < color ||
					int(scale(world.Terrain[x][chompInt(y-1, 0, world.Height-1)], -1.0, 1.0, 0.0, 31.0)) < color ||
					int(scale(world.Terrain[chompInt(x-1, 0, world.Width-1)][y], -1.0, 1.0, 0.0, 31.0)) < color {
					color = black
				} else {
					color = white
				}

				if int(scale(world.Terrain[chompInt(x+1, 0, world.Width-1)][y], -1.0, 1.0, 0.0, 31.0)) == sealevel-1 && (y%4 == 0) {
					color = wave
				}

				world.Map.Set(x, y, palette[color])
			}
		}
	}

}

// returns an array of colors from a .png image
func createPalette(path string) (palette color.Palette) {
	file, _ := os.Open(path)
	defer file.Close()

	imagePalette, _, _ := image.Decode(file)

	for i := 0; i < 32; i++ {
		palette = append(palette, imagePalette.At(i, 0))
	}

	return
}

func chompInt(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

func chompFloat(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

// transforms a number in one range to another range
func scale(value, oldMin, oldMax, newMin, newMax float64) float64 {
	return (value-oldMin)*(newMax-newMin)/(oldMax-oldMin) + newMin
}

// saves the current world's map to disk
func (world *World) exportMap() {
	file, err := os.Create("out/map.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, &world.Map)
}

// (TODO) generates a new random name for the world
func (world *World) name() {
	world.Name = "New World"
}
