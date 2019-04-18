package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type World struct {
	Terrain [][]float64 // generated 2d grid for the world's elevation
	Map     image.RGBA  // generated map from elevation and a mapview
	Name    string      // generated random name for the world
	Width   int         // world width
	Height  int         // world height
}

// MapView determines aspects like color palettes
// and how the map is generated from the elevation grid
type MapView uint8

const (
	ElevationView        MapView = iota                         // default view, represents the terrain elevation
	ClimateView                                                 // climate map view, represents an average temperature
	PoliticalView                                               // colored regions representing political states that would form based on the terrain
	BiomeView                                                   // biome based on climate, elevation, wet/dryness
	AssetsDir            string  = "assets/"                    // path to the assets directory
	PaletteDir           string  = AssetsDir + "palettes/"      // path to the palette directory
	ElevationPalettePath string  = PaletteDir + "elevation.png" // path to the elevation palette .png
	ClimatePalettePath   string  = PaletteDir + "climate.png"   // path to the climatepalette .png
	PoliticalPalettePath string  = PaletteDir + "political.png" // path to the political palette .png
	BiomePalettePath     string  = PaletteDir + "biome.png"     // path to the biome palette .png
)

// generates a completely new world from scratch
func newWorld(width, height int) (world *World) {
	// create a new, blank world
	world = new(World)

	// give the world dimensions
	world.Width = width
	world.Height = height

	// initialize the elevation height map with 0s
	world.Terrain = make([][]float64, world.Width)
	for x := range world.Terrain {
		world.Terrain[x] = make([]float64, world.Height)
	}

	world.generateTerrain()
	world.drawMap(ElevationView) // elevation is the default mapview
	world.name()

	return
}

// generates a height map from scratch for a new world
func (world *World) generateTerrain() {
	world.Terrain = addOpenSimplexNoise(world.Terrain, 0, world.Width, 0, world.Height)
}

// creates a map image based on the world and provided mapview
func (world *World) drawMap(mapView MapView) {
	// create a new, blank map image
	world.Map = *image.NewRGBA(image.Rect(0, 0, world.Width, world.Height))

	// load the palette for the selected map view
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
	}

	// write the map image
	switch mapView {
	case ElevationView:
		for x := 0; x < world.Map.Bounds().Max.X; x++ {
			for y := 0; y < world.Map.Bounds().Max.Y; y++ {
				// map the grid value to a color in the palette
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))

				// write the color to the image
				world.Map.Set(x, y, palette[i])
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
		for x := 0; x < world.Map.Bounds().Max.X; x++ {
			for y := 0; y < world.Map.Bounds().Max.Y; y++ {
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	case BiomeView: // (TODO) alogrithm for interpreting biome based on climate, elevation, proximity to ocean, etc.
		for x := 0; x < world.Map.Bounds().Max.X; x++ {
			for y := 0; y < world.Map.Bounds().Max.Y; y++ {
				i := int(scale(world.Terrain[x][y], -1.0, 1.0, 0.0, 31.0))
				world.Map.Set(x, y, palette[i])
			}
		}
	}
}

// returns an array of colors from a .png image
func createPalette(path string) (palette color.Palette) {
	// open the file at the path
	file, _ := os.Open(path)
	defer file.Close()

	// decode to a go-library image.RGBA
	imagePalette, _, _ := image.Decode(file)

	// read the images pixel into a color palette
	for i := 0; i < 31; i++ {
		palette = append(palette, imagePalette.At(i, 0))
	}

	return
}

// transforms a number in one range to another range
func scale(value, oldMin, oldMax, newMin, newMax float64) float64 {
	return (value-oldMin)*(newMax-newMin)/(oldMax-oldMin) + newMin
}

// saves the current world's map to disk
func (world *World) exportMap() {
	// open and prepare a new file
	file, err := os.Create("out/map.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// encode and write a .png from the world map
	png.Encode(file, &world.Map)
}

// (TODO) generates a new random name for the world
func (world *World) name() {
	world.Name = "New World"
}
