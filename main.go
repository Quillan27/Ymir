package main

import (
	"fmt"
)

const ELEVATION int = 0
const CLIMATE int = 1
const POLITICAL int = 2
const BIOME int = 3

func main() {
	fmt.Printf("Welcome to Ymir!\n")
	fmt.Printf("----------------\n\n")

	w := newWorld(27, 27)
	w.exportMap()
}
