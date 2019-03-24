package main 

import (
    "fmt"
    "image"
    "image/color"
    "bufio"
    "os"
)

func (w *World) updateMap(view int) {
    p := createPalette(view)
    w.Image = *image.NewPaletted(image.Rect(0, 0, len(w.Grid[view]), len(w.Grid[view][0])), p)
    for x := 0; x < w.Image.Bounds().Max.X; x++ {
	for y := 0; y < w.Image.Bounds().Max.Y; y++ {
	    c = p[w.Grid[view][x][y]]
	    w.Image.Set(x, y, c) 
	}
    }
}

func (w *World) saveMap() {

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
  
    hexColors, err := readLines(path)
    if err != nil {
	fmt.Print(path + "could not be opened.\n")
    }
    
    fmt.Print(hexColors)
    return *new(color.Palette)
} 

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
