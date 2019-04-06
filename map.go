package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

type MapView int

const (
	ELEVATION MapView = iota
	CLIMATE   MapView = iota
	POLITICAL MapView = iota
	BIOME     MapView = iota
)

func (world *World) drawMap(mapView MapView) {
	world.Map = *image.NewRGBA(image.Rect(0, 0, len(world.Elevation), len(world.Elevation[0])))

	palette := color.Palette{}
	switch mapView {
	case ELEVATION:
		palette = createPalette("assets/colors/elevation.txt")
		for x := 0; x < world.Map.Bounds().Max.X; x++ {
			for y := 0; y < world.Map.Bounds().Max.Y; y++ {
				i := int((world.Elevation[x][y] - -0.707)*(32.0-0.0)/(0.707 - -0.707) + 0.0)
				world.Map.Set(x, y, palette[i])
			}
		}
	case CLIMATE:
		palette = createPalette("assets/colors/climate.txt")
	case POLITICAL:
		palette = createPalette("assets/colors/poltical.txt")
	case BIOME:
		palette = createPalette("assets/colors/biome.txt")
	}
}

func createPalette(path string) (palette color.Palette) {
	hexColors, err := splitLines(path)
	if err != nil {
		fmt.Print(path + "could not be opened.\n")
	}

	for i := 0; i < len(hexColors); i++ {
		red, _ := strconv.ParseUint(hexColors[i][1:3], 16, 8)
		green, _ := strconv.ParseUint(hexColors[i][3:5], 16, 8)
		blue, _ := strconv.ParseUint(hexColors[i][5:7], 16, 8)
		alpha := 255
		color := color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alpha)}
		palette = append(palette, color)
	}

	return
}

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

func (world *World) exportMap() {
	file, err := os.Create("map.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, &world.Map)
}
