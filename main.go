package main

import (
	"fmt"
)

const ELEVATION int = 0
const CLIMATE int = 1
const POLITICAL int = 2
const BIOME int = 3

func main() {
    fmt.Printf("Welcome to Ymir!\n\n")

    w := newWorld(10, 10)	
    
    fmt.Print(w.Name+"\n")
    fmt.Print("-----------\n")
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
		fmt.Printf("%02d", int(w.Grid[i][j][k] * 31.0))
		fmt.Printf(" ")
	    }
	    fmt.Printf("\n")
	}
	fmt.Printf("\n")
    }
    w.saveMap()
}
