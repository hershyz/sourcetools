package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: varvals (path)")
		return
	}

	lines := getFileLines(args[0])
	var i int = 0
	for i < len(lines) {
		if containsVar(lines[i]) {
			varVal := getVarValue(lines[i])
			if varVal == "" {
				fmt.Println(lines[i])
			}
			if len(varVal) > 0 {
				name := strings.Split(lines[i], "=")[0]
				name = strings.TrimSpace(name)
				fmt.Print(name + ", ")
				fmt.Println(varVal)
			}
		}
		i++
	}
}
