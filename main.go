package main

import (
	"fmt"
	"math/rand"
)

type World struct {
    Grid [][][]float32
    Name string
}
   
func newWorld() (w *World) {
    w = &World{}
    
    w.Grid = make([][][]float32, 4)
    for i := range w.Grid {
	w.Grid[i] = make([][]float32, 10)
	for j := range w.Grid[i] {
	    w.Grid[i][j] = make([]float32, 10)
	}
    }
    
    genWorld(w)

    w.Name = newWorldName()

    return w
}

func genWorld(w *World) {
    for i := range w.Grid {
	for j := range w.Grid[i] {
	    for k := range w.Grid[i][j] {
		w.Grid[i][j][k] += float32(rand.Intn(256))
	    }
	}
    }
}

func newWorldName() string {
    return "New World"
}

func main() {
    fmt.Printf("Welcome to Ymir!\n")
    w := *newWorld()	
    for i := range w.Grid {
	for j := range w.Grid[i] {
	    fmt.Println(w.Grid[i][j])
	}
	fmt.Printf("\n")
    }
}
