package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"bufio"
	"os"
	"strconv"
)

func (w *World) drawMap(view int) {
    fmt.Printf("Drawing the map...\n")

    // get color palette for map view
    p := createPalette(view)

    // make new image based on grid
    w.Map = *image.NewRGBA(image.Rect(0, 0, len(w.Grid[view]), len(w.Grid[view][0])))
    for y := 0; y < w.Map.Bounds().Max.X; y++ {
	for x := 0; x < w.Map.Bounds().Max.Y; x++ {
	    c := p[int(w.Grid[x][y][view]*32.0)]
	    w.Map.Set(x, y, c)
	}
    }
}

func createPalette(view int) (p color.Palette) {
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

func (w *World) saveMap() {
    fmt.Printf("Saving the map to disk...\n")

    f, err := os.Create("map.png")
    if err != nil {
	panic(err)
    }
    defer f.Close()

    png.Encode(f, &w.Map)
}

