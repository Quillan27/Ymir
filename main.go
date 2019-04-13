package main

import (
	"os"
	"strconv"
)

func main() {
	var width, height int
	if os.Args != nil {
		args := os.Args[1:]
		width, _ = strconv.Atoi(args[0])
		height, _ = strconv.Atoi(args[1])
	} else {
		width = 512
		height = 512
	}

	world := newWorld(width, height)
	world.exportMap()
	startServer()
}
