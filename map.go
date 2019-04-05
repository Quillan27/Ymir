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

func (w *World) drawMap(view int) {
	w.Map = *image.NewRGBA(image.Rect(0, 0, len(w.Elevation), len(w.Elevation[0])))

	p := color.Palette{}
	switch view {
	case ELEVATION:
		p = createPalette("res/colors/bandw.txt")
		for x := 0; x < w.Map.Bounds().Max.X; x++ {
			for y := 0; y < w.Map.Bounds().Max.Y; y++ {
				i := int((w.Elevation[x][y] - -0.707)*(32.0-0.0)/(0.707 - -0.707) + 0.0)
				w.Map.Set(x, y, p[i])
			}
		}
	case CLIMATE:
		p = createPalette("res/colors/climate.txt")
	case POLITICAL:
		p = createPalette("res/colors/poltical.txt")
	case BIOME:
		p = createPalette("res/colors/biome.txt")
	}
}

func createPalette(path string) (p color.Palette) {
	hexColors, err := splitLines(path)
	if err != nil {
		fmt.Print(path + "could not be opened.\n")
	}

	for i := 0; i < len(hexColors); i++ {
		r, _ := strconv.ParseUint(hexColors[i][1:3], 16, 8)
		g, _ := strconv.ParseUint(hexColors[i][3:5], 16, 8)
		b, _ := strconv.ParseUint(hexColors[i][5:7], 16, 8)
		var a uint8 = 255
		c := color.RGBA{uint8(r), uint8(g), uint8(b), a}
		p = append(p, c)
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

func (w *World) exportMap() {
	f, err := os.Create("map.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, &w.Map)
}
