package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"
)

// World holds all the basic world data
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
	/*** MapView ENUMERATION ***/

	// ElevationView shows terrain elevation
	ElevationView MapView = iota

	// ClimateView shows average yearly temperatures
	ClimateView

	// PoliticalView shows political boundries
	PoliticalView

	// BiomeView shows biomes (duh)
	// TODO(karl): see BiomeType enum for all the types
	BiomeView

	// TopographyView shows elevation through topography levels
	TopographyView

	/*** ASSET FILE PATHS ***/

	// AssetDir is the path to program assets
	AssetDir string = "assets/"

	// PaletteDir is the path to MapView palettes
	PaletteDir string = AssetDir + "palettes/"

	// ElevationPalettePath is the path to the ElevationView palette
	ElevationPalettePath string = PaletteDir + "elevation.png"

	// BiomePalettePath is the path to the BiomeView palette
	BiomePalettePath string = PaletteDir + "biome.png"

	// PoliticalPalettePath is the path to the PoliticalView palette
	PoliticalPalettePath string = PaletteDir + "political.png"

	// ClimatePalettePath is the path to the ClimateView palette
	ClimatePalettePath string = PaletteDir + "climate.png"

	// TopographyPalettePath is the path to the TopographyView palette
	TopographyPalettePath string = PaletteDir + "topography.png"
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
	world.drawMap(ElevationView) // ElevatioView is the default
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
	case BiomeView:
		palette = createPalette(BiomePalettePath)
	case PoliticalView:
		palette = createPalette(PoliticalPalettePath)
	case ClimateView:
		palette = createPalette(ClimatePalettePath)
	case TopographyView:
		palette = createPalette(TopographyPalettePath)
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
	case BiomeView: // (TODO) algorithm for interpreting elevation for climate
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
	case ClimateView: // (TODO) alogrithm for interpreting biome based on climate, elevation, proximity to ocean, etc.
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	case TopographyView: // black if next to a lower elevation, white if below sealevel or otherwise
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

// reads in a file and return the lines as an slice of strings
func splitLines(path string) (lines []string, scanErr error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	scanErr = scanner.Err()

	return
}

// (TODO) generates a new random name for the world
func (world *World) name() {
	syllables, err := splitLines("assets/naming/world.txt")
	if err != nil {
		fmt.Print(err)
	}

	var name string
	rand.Seed(time.Now().UnixNano())
	numOfSyllables := int(2 + rand.Float64()*2)
	for i := 0; i < numOfSyllables; i++ {
		name += syllables[int(rand.Float64()*float64(len(syllables)-1))]
	}

	name = strings.Title(name)

	fmt.Print(name, "\n")
	world.Name = name
}
