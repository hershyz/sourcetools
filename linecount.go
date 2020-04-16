package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: linecount (path)")
		return
	}

	lines := getFileLines(args[0])
	fmt.Println(len(lines))
}
