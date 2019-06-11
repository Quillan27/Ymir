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

type World struct {
	Name     string
	Terrain  [][]uint8
	Biomes   [][]BiomeID 
	Climate  [][]uint8
	Width    int
	Height   int
	Map      image.RGBA
}

type MapView uint8

const (
	ElevationView MapView = iota
	TopographyView
	BiomeView
	ClimateView
)

const (
	MinElev     uint8 = 0
	MaxElev     uint8 = 255
	Octaves     int     = 8
	Persistence float64 = 2.0
)

const (
	LevelChange uint8 = iota
	Flat
	Shading
	Water
)

type BiomeID uint8

const (
	// TODO(karl): add BiomeID's here
)

var currentView = ElevationView

func newWorld(width, height int) *World {
	w := new(World)
	w.Width = width
	w.Height = height

	w.Terrain = make([][]uint8, w.Width)
	for x := range w.Terrain {
		w.Terrain[x] = make([]uint8, w.Height)
	}

	w.Climate = make([][]uint8, w.Width)
	for x := range w.Climate {
		w.Climate[x] = make([]uint8, w.Height)
	}

	w.Biomes = make([][]BiomeID, w.Width)
	for x := range w.Biomes {
		w.Biomes[x] = make([]BiomeID, w.Height)
	}

	w.generate()
	w.addName()
	w.drawMap(currentView)
	return w
}

// TEMPORARY - for testing algorithms
func (w *World) drawMap(mapView MapView) {
	w.Map = *image.NewRGBA(image.Rect(0, 0, w.Width, w.Height))
	palette := createPalette("assets/palettes/test.png")
	for x := range w.Terrain {
		for y:= range w.Terrain[x] {
			w.Map.Set(x, y, palette[w.Terrain[x][y]])
		}
	}
	currentView = mapView
}
/*
func (w *World) drawMap(mapView MapView) {
	w.Map = *image.NewRGBA(image.Rect(0, 0, w.Width, w.Height))
	palette := colkor.Palette{}
	switch mapView {
	case ElevationView:
		palette = createPalette("assets/palettes/elevation.png")

		var color uint8
		for x := range w.Terrain {
			for y := range w.Terrain[x] {
				color = calculateTopographicLevel(w.Terrain[x][y], len(palette)-1)
				w.Map.Set(x, y, palette[color])
			}
		}
	case TopographyView:
		palette = createPalette("assets/palettes/topography.png")

		const MaxLevel uint8 = 31 // I found 31 to produce the best results
		seaLevel := uint8(MaxLevel / 2)

		var color uint8
		for x := range w.Terrain {
			for y := range w.Terrain[x] {
				// map elevation to a topographic level (0 - 31)
				level := calculateTopographicLevel(w.Terrain[x][y], MaxLevel)

				// look at surrounding levels
				up := calculateTopographicLevel(w.Terrain[x][min(y+1, w.Height-1)], MaxLevel)
				down := calculateTopographicLevel(w.Terrain[x][max(y-1, 0)], MaxLevel)
				left := calculateTopographicLevel(w.Terrain[max(x-1, 0)][y], MaxLevel)
				right := calculateTopographicLevel(w.Terrain[min(x+1, w.Width-1)][y], MaxLevel)
				color = Flat
				if level <= seaLevel {
					color = Water
				} else if level > up || level > down || level > left || level > right {
					color = LevelChange
				}

				if level == seaLevel && y%4 == 0 {
					color = Shading
				}
				w.Map.Set(x, y, palette[color])
			}
		}
	case BiomeView:
		palette = createPalette("assets/palettes/biome.png")

		var color uint8
		for x := range w.Biomes {
			for y := range w.Biomes[x] {
				color = uint8(w.Biomes[x][y])
				w.Map.Set(x, y, palette[color])
			}
		}
	case ClimateView:
		palette = createPalette("assets/palettes/climate.png")

		var color uint8
		for x := range w.Climate {
			for y := range w.Climate[x] {
				color = w.Climate[x][y]
				w.Map.Set(x, y, palette[color])
			}
		}
	}

	currentView = mapView
}
*/
func createPalette(path string) color.Palette {
	file, _ := os.Open(path)
	defer file.Close()

	img, _, _ := image.Decode(file)

	var palette color.Palette
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			palette = append(palette, img.At(x, y))
		}
	}

	return palette
}
/*
func calculateTopographicLevel(elevation uint8, maxLevel uint8) uint8 {
	return int((elevation - MinElev) * float64(maxLevel) / (MaxElev - MinElev))
} */

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (w *World) exportMap() {
	file, err := os.Create("out/map.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, &w.Map)
}

func (w *World) addName() {
	const MinWordLength = 2
	const MaxWordLength = 4

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

func splitLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	return lines, err
}
