package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("usage: parameters (path, function name)")
		return
	}

	lines := getFileLines(args[0])
	funcName := args[1]
	var i int = 0

	for i < len(lines) {

		if strings.Contains(lines[i], funcName) {
			if isFunction(lines[i]) {
				fmt.Println(getRawParameters(lines[i]))
			}
		}

		i++
	}
}
