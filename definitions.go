package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("usage: definitions (path, term)")
		return
	}

	filepath := args[0]
	term := args[1]

	lines := getFileLines(filepath)
	var i int = 0
	for i < len(lines) {
		if strings.Contains(lines[i], term) {
			if containsTrigger(lines[i]) {
				lineNum := i + 1
				fmt.Println("line ", lineNum, " - ", lines[i])
			}
		}
		i++
	}
}
