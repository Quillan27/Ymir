package main

import (
	"fmt"
	"os"
	"strconv"
)

const ELEVATION int = 0
const CLIMATE int = 1
const POLITICAL int = 2
const BIOME int = 3

func main() {
	fmt.Printf("Welcome to Ymir!\n")
	fmt.Printf("----------------\n\n")

	args := os.Args[1:]
	width, _ := strconv.Atoi(args[0])
	height, _ := strconv.Atoi(args[1])
	w := newWorld(width, height)
	w.exportMap()
}
