package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: compress (path)")
		return
	}

	lines := getFileLines(args[0])
	var i int = 0
	var final string = ""

	for i < len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
		final = final + lines[i]
		i++
	}

	fmt.Println(final)
}
