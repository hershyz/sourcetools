package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: read (path)")
		return
	}

	var i int = 0
	lines := getFileLines(args[0])
	for i < len(lines) {
		fmt.Println(lines[i])
		i++
	}
}
