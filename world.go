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
	Name     string
	Terrain  [][]float64
	Biomes   [][]uint8
	Climate  [][]uint8
	Politics [][]uint8
	Width    int
	Height   int
	Map      image.RGBA
}

// MapView determines what color palette and grid to use for the world's Map
type MapView uint8

const (
	// ElevationView shows terrain elevation
	ElevationView MapView = iota
	// TopographyView shows terrain with levels
	TopographyView
	// BiomeView shows biomes
	BiomeView
	// PoliticalView shows political boundries
	PoliticalView
	// ClimateView shows average yearly temperatures
	ClimateView
)

const (
	// MinElev is the minimum elevation possible
	MinElev int = -1
	// MaxElev is the maximum elevation possible
	MaxElev int = 1
	// Octaves controls how many times the noise function is called
	Octaves int = 8
	// Persistence controls how much effect the noise will have over time
	Persistence float64 = 2.0
)

const (
	// LevelChange is the index for the color used to mark a change
	// the terrain's topographic level
	LevelChange uint8 = iota
	// Flat is the index for the color used when there is no different
	// topographic level for any surrounding points
	Flat
	// Shading is the index for color used for aesthetics, seperating water from land
	Shading
	// Water is the index for the color used for any point at or below SeaLevel
	Water
	// LevelCount is the number of topographic levels the terrain is split into
	// I found 31 to be the sweet spot to see general trends in terrain as well as detail
	LevelCount int = 31
)

var currentView = ElevationView

// newWorld acts like a constructor for the World struct
// it initializes terrain, names, and map
func newWorld(width, height int) *World {
	w := new(World)

	w.Width = width
	w.Height = height

	w.generateTerrain()
	w.name()
	w.drawMap(currentView)

	return w
}

// generateTerrain initializes the terrain's 2D array and generates elevations
func (w *World) generateTerrain() {
	w.Terrain = make([][]float64, w.Width)
	for x := range w.Terrain {
		w.Terrain[x] = make([]float64, w.Height)
	}

	addPerlinNoise(&w.Terrain, Octaves, Persistence)

	for x := range w.Terrain {
		for y, v := range w.Terrain[x] {
			// limit noise to a range between MinElev and MaxElev
			if v > float64(MaxElev) {
				w.Terrain[x][y] = float64(MaxElev)
			} else if v < float64(MinElev) {
				w.Terrain[x][y] = float64(MinElev)
			}
		}
	}
}

// drawMap draws a new map based on the world's information and the passed MapView
func (w *World) drawMap(mapView MapView) {
	w.Map = *image.NewRGBA(image.Rect(0, 0, w.Width, w.Height))
	currentView = mapView

	palette := color.Palette{}
	switch mapView {
	case ElevationView:
		palette = createPalette("assets/palettes/elevationPalette.png")
	case TopographyView:
		palette = createPalette("assets/palettes/topographyPalette.png")
	case BiomeView:
		palette = createPalette("assets/palettes/biomePalette.png")
	case ClimateView:
		palette = createPalette("assets/palettes/climatePalette.png")
	case PoliticalView:
		palette = createPalette("assets/palettes/politicalPalette.png")
	}

	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			var color uint8
			switch mapView {
			case ElevationView:
				color = uint8(scale(w.Terrain[x][y], MinElev, MaxElev, 0, len(palette)-1))
			case TopographyView:
				seaLevel := LevelCount / 2
				// map elevation to a topographic level (0 - 31)
				level := int(scale(w.Terrain[x][y], MinElev, MaxElev, 0.0, LevelCount))

				// look at surrounding levels
				up := int(scale(w.Terrain[x][chomp(y+1, 0, w.Height-1)], MinElev, MaxElev, 0.0, LevelCount))
				down := int(scale(w.Terrain[x][chomp(y-1, 0, w.Height-1)], MinElev, MaxElev, 0.0, LevelCount))
				left := int(scale(w.Terrain[chomp(x-1, 0, w.Width-1)][y], MinElev, MaxElev, 0.0, LevelCount))
				right := int(scale(w.Terrain[chomp(x+1, 0, w.Width-1)][y], MinElev, MaxElev, 0.0, LevelCount))

				color = Flat
				if level <= seaLevel {
					color = Water
				} else if level > up || level > down || level > left || level > right {
					color = LevelChange
				}

				if level == seaLevel && y%4 == 0 {
					color = Shading
				}
			case BiomeView:
				color = w.Biomes[x][y]
			case ClimateView:
				color = w.Climate[x][y]
			case PoliticalView:
				color = w.Politics[x][y]
			}
			w.Map.Set(x, y, palette[color])
		}
	}

}

// createPalette returns an array of colors from a premade .png image
func createPalette(path string) (p color.Palette) {
	f, _ := os.Open(path)
	defer f.Close()

	i, _, _ := image.Decode(f)

	for x := 0; x < i.Bounds().Max.X; x++ {
		for y := 0; y < i.Bounds().Max.Y; y++ {
			p = append(p, i.At(x, y))
		}
	}

	return
}

// scale maps a number in one range to another range
func scale(value float64, oldMin, oldMax, newMin, newMax int) float64 {
	return (value-float64(oldMin))*float64(newMax-newMin)/float64(oldMax-oldMin) + float64(newMin)
}

// chompInt keeps an int inside a specified range
func chomp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

// exportMap exports the current world's map to disk
// TODO(karl): create an actual save window popup
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

// name names the world using a random sequence of pre-defined sylablles
// TODO(karl): clean up syllables so names are pronouncable
// BUG(karl): name seems to be one-behind
func (w *World) name() {
	const MinWordLength = 2
	const MaxWordLength = 4

	// TODO(karl): filename constant
	syllables, err := splitLines("assets/syllables/world.txt")
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	var name string
	wordLength := MinWordLength + int(rand.Float64()*float64(MaxWordLength-MinWordLength))
	for i := 0; i < wordLength; i++ {
		name += syllables[int(rand.Float64()*float64(len(syllables)-1))]
	}

	w.Name = strings.Title(name)
}
