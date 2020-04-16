package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("usage: copy (existing path, new path)")
		return
	}

	original := args[0]
	new := args[1]

	lines := getFileLines(original)
	createFile(new)
	var i int = 0

	for i < len(lines) {
		writeLineToFile(new, lines[i])
		i++
	}
}
