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

// redraws map based on the world grid
func (w *World) updateMap(view int) {
	// get color palette for map view
	p, hexes := createPalette(view)

	// make new image based on grid
	w.Map = *image.NewRGBA(image.Rect(0, 0, len(w.Grid[view]), len(w.Grid[view][0])))
	for y := 0; y < w.Map.Bounds().Max.X; y++ {
		for x := 0; x < w.Map.Bounds().Max.Y; x++ {
			c := p[int(w.Grid[x][y][view]*32.0)]
			w.Map.Set(x, y, c)
		}
	}
}

// saves a map.png to the disk
func (w *World) saveMap() {
	// create a new file
	f, err := os.Create("map.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write in png format
	png.Encode(f, &w.Map)
}

// makes a array of colors from a text file
func createPalette(view int) (color.Palette, []string) {
	// get colors from a file in hexidecimal format
	var path string
	if view == ELEVATION {
		path = "res/palettes/elevation.pal"
	} else if view == CLIMATE {
		path = "res/palettes/climate.pal"
	} else if view == POLITICAL {
		path = "res/palettes/political.pal"
	} else {
		path = "res/palettes/biome.pal"
	}
	
	hexColors, err := readLines(path)
	if err != nil {
		fmt.Print(path + "could not be opened.\n")
	}

	// convert hexidecimal colors to RGBA and add to a palette
	p := color.Palette{}
	for i := 0; i < len(hexColors); i++ {
		r, _ := strconv.ParseUint(hexColors[i][1:3], 16, 8)
		g, _ := strconv.ParseUint(hexColors[i][3:5], 16, 8)
		b, _ := strconv.ParseUint(hexColors[i][5:7], 16, 8)
		var a uint8 = 255
		c := color.RGBA{uint8(r), uint8(g), uint8(b), a}
		p = append(p, c)
	}

	return p, hexColors
}

// opens a file and splits it's line into an array
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
