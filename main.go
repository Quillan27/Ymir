package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	width, _ := strconv.Atoi(args[0])
	height, _ := strconv.Atoi(args[1])
	w := newWorld(width, height)
	w.exportMap()
}
