package main

import (
	"bufio"
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
	Width   int // used for readability instead of using
	Height  int // len(Terrain) or len(Terrain[0])
	Map     image.RGBA
}

// MapView determines aspects like color palettes
// and how the map is generated from the terrain
type MapView uint8

// NoiseType determines the algorithms for world generation
// for now only Perlin noise and OpenSimplex are implemented
// TODO(karl): set noise based on user settings
type NoiseType uint8

// BiomeType is used to enumerate biomes
// TOOD(karl): make this enumeration
type BiomeType uint8

const (
	// MAPVIEW ENUMERATION -----------------------------------------------------

	// ElevationView shows terrain elevation
	ElevationView MapView = iota

	// ClimateView shows average yearly temperatures
	ClimateView

	// PoliticalView shows political boundries
	PoliticalView

	// BiomeView shows biomes (duh)
	// TODO(karl): make BiomeType enum
	BiomeView

	// TopographyView shows elevation through topography levels
	TopographyView

	// ASSET FILE PATHS --------------------------------------------------------

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

	// NOISE CONSTANTS ---------------------------------------------------------

	// Octaves controls how many times the noise function is called
	Octaves int = 8

	// Persistence controls how much effect the noise will have over time
	Persistence float64 = 2.0

	// NAMING CONSTANTS --------------------------------------------------------

	// MinSyllables is the minimum number of syllables a name can have
	MinSyllables int = 2

	// MaxSyllables is the maximum number of syllables a name can have
	MaxSyllables int = 4
)

// newWorld starts acts like a constructor for the World struct
// it initializes terrain, names, and map
func newWorld(width, height int) (w *World) {
	w = new(World)

	w.Width = width
	w.Height = height

	w.generateTerrain()

	// TODO(karl): set to current view, for now ElevationView is default
	w.drawMap(ElevationView)

	w.name()

	return
}

// generateTerrain initializes the terrain's 2D array and generates elevations
// TODO(karl): more realistic / complex than just adding noise,
// generation should be based on plate tectonics and other
// realistic systems
func (w *World) generateTerrain() {
	w.Terrain = make([][]float64, w.Width)
	for x := range w.Terrain {
		w.Terrain[x] = make([]float64, w.Height)
	}

	addPerlinNoise(&w.Terrain, Octaves, Persistence)
}

// drawMap draws a new map based on the world's terrain and the passed MapView
func (w *World) drawMap(mapView MapView) {
	// create a new, blank RGBA image
	w.Map = *image.NewRGBA(image.Rect(0, 0, w.Width, w.Height))

	// get the map's color palette based on the MapView
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

	// interpret colors based on the MapView
	// TODO(karl): clean up magic numbers here
	switch mapView {
	case ElevationView:
		for x := 0; x < w.Width; x++ {
			for y := 0; y < w.Height; y++ {
				// ensure the elevation is between -1.0 and 1.0
				w.Terrain[x][y] = chompFloat(w.Terrain[x][y], -1.0, 1.0)

				// map the grid value to a color in the palette between 0 and 31
				color := int(scale(w.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))

				w.Map.Set(x, y, palette[color])
			}
		}
	// TODO(karl): algorithm for interpreting biome based on elevation, climate,
	// and moisture
	case BiomeView:
		for x := 0; x < w.Map.Bounds().Max.X; x++ {
			for y := 0; y < w.Map.Bounds().Max.Y; y++ {
				i := int(scale(w.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				w.Map.Set(x, y, palette[i])
			}
		}
	// TODO(karl): algorithm for interpreting political boundaries based on
	// terrain
	case PoliticalView:
		for x := 0; x < w.Width; x++ {
			for y := 0; y < w.Height; y++ {
				i := int(scale(w.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				w.Map.Set(x, y, palette[i])
			}
		}
	// TODO(karl): alogrithm for interpreting climate
	case ClimateView:
		for x := 0; x < w.Width; x++ {
			for y := 0; y < w.Height; y++ {
				i := int(scale(w.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				w.Map.Set(x, y, palette[i])
			}
		}
	// black if next to a lower elevation, white if below sealevel or otherwise
	case TopographyView:
		for x := 0; x < w.Width; x++ {
			for y := 0; y < w.Height; y++ {
				w.Terrain[x][y] = chompFloat(w.Terrain[x][y], -1.0, 1.0)

				color := int(scale(w.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))

				const Black int = 0
				const White int = 31
				const Water int = 27
				const Wave int = 17
				const SeaLevel int = 16
				if color < SeaLevel {
					color = Water
				} else if int(scale(w.Terrain[chompInt(x+1, 0, w.Width-1)][y], -1.0, 1.0, 0.0, 31.0)) < color ||
					int(scale(w.Terrain[x][chompInt(y+1, 0, w.Height-1)], -1.0, 1.0, 0.0, 31.0)) < color ||
					int(scale(w.Terrain[x][chompInt(y-1, 0, w.Height-1)], -1.0, 1.0, 0.0, 31.0)) < color ||
					int(scale(w.Terrain[chompInt(x-1, 0, w.Width-1)][y], -1.0, 1.0, 0.0, 31.0)) < color {
					color = Black
				} else {
					color = White
				}

				if int(scale(w.Terrain[chompInt(x+1, 0, w.Width-1)][y], -1.0, 1.0, 0.0, 31.0)) == SeaLevel-1 && (y%4 == 0) {
					color = Wave
				}

				w.Map.Set(x, y, palette[color])
			}
		}
	}
	w.exportMap() // TOD0(karl): remove this, map to a button
}

// createPalette returns an array of colors from a premade .png image
func createPalette(path string) (p color.Palette) {
	f, _ := os.Open(path)
	defer f.Close()

	i, _, _ := image.Decode(f)

	// TODO(karl): 32 needs to be replaces by the actual number of colors
	for x := 0; x < 32; x++ {
		p = append(p, i.At(x, 0))
	}

	return
}

// chompInt keeps an int inside a specified range
// TODO(karl): go needs generics, lol
func chompInt(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

// chompFloat keeps a float64 inside a specified range
// TODO(karl): i hate that these are two functions
func chompFloat(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

// scale maps a number in one range to another range
func scale(value, oldMin, oldMax, newMin, newMax float64) float64 {
	return (value-oldMin)*(newMax-newMin)/(oldMax-oldMin) + newMin
}

// exportMap exports the current world's map to disk
func (w *World) exportMap() {
	f, err := os.Create("out/map.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, &w.Map)
}

// splitLines reads in a file and return the lines as a slice of strings
func splitLines(path string) (lines []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	err = s.Err()

	return
}

// name names the world
// generates using a random sequence of pre-defined sylablles
// TODO(karl): clean up syllables so names are pronouncable
// BUG(karl): name seems to be one-behind
func (w *World) name() {
	// TODO(karl): filename constant
	syllables, err := splitLines("assets/naming/world.txt")
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	var name string
	// TODO(karl): can this be cleaner without casting?
	numOfSyllables := MinSyllables + int(rand.Float64()*float64(MaxSyllables-MinSyllables))
	for i := 0; i < numOfSyllables; i++ {
		// TODO(karl): again, can we clean this up?
		name += syllables[int(rand.Float64()*float64(len(syllables)-1))]
	}

	w.Name = strings.Title(name)
}
