package main

import (
	"fmt"
	"math/rand"
)

const ELEVATION int = 0
const CLIMATE int = 1
const POLITICAL int = 2
const BIOME int = 3

type World struct {
    Grid [][][]float64
    Name string
}
  
func newWorld(width, height int) (*World) {
    w := new(World)
    
    w.Grid = make([][][]float64, 4)
    for i := range w.Grid {
	w.Grid[i] = make([][]float64, width)
	for j := range w.Grid[i] {
	    w.Grid[i][j] = make([]float64, height)
	}
    }
    genWorld(w)
    nameWorld(w)

    return w
}

func genWorld(w *World) {
    for i := range w.Grid {
	for j := range w.Grid[i] {
	    for k := range w.Grid[i][j] {
		w.Grid[i][j][k] += float64(rand.Float64())
	    }
	}
    }
}

func nameWorld(w *World) {
    w.Name = "New World"
}

func main() {
    fmt.Printf("Welcome to Ymir!\n")
    w := *newWorld(10, 10)	
    for i := range w.Grid {
	if i == ELEVATION {
	    fmt.Print("Elevation\n")
	} else if i == CLIMATE {
	    fmt.Print("Climate\n")
	} else if i == POLITICAL {
	    fmt.Print("Political\n")
	} else {
	    fmt.Print("Biome\n")
	}

	for j := range w.Grid[i] {
	    for k := range w.Grid[i][j] {
		fmt.Printf("%02d", int(w.Grid[i][j][k]*31.0))
		fmt.Printf(" ")
	    }
	    fmt.Printf("\n")
	}
	fmt.Printf("\n")
    }
}
