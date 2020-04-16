package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: nocomment (path)")
		return
	}

	lines := getFileLines(args[0])
	var i int = 0
	for i < len(lines) {
		if len(lines[i]) > 0 {
			if !startsWithComment(lines[i]) {
				fmt.Println(lines[i])
			}
		}

		i++
	}
}
