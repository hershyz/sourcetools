package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: structure (path)")
		return
	}

	lines := getFileLines(args[0])
	var i int = 0
	var insideVoid = false

	for i < len(lines) {

		if strings.Contains(lines[i], "}") {
			insideVoid = false
		}

		if insideVoid {
			if containsVar(lines[i]) {
				fmt.Println(lines[i])
			}
		}

		if strings.Contains(lines[i], "{") {
			insideVoid = true
			fmt.Println(lines[i])
		}

		i++
	}
}
