package main

import (
	"bufio"
	"os"
	"strings"
)

var triggers = []string{"void", "def", "function", "func", "int", "var", "double", "float", "string", "String", "bool", "boolean"}

func getFileLines(filepath string) []string {
	currentfile, err := os.Open(filepath)

	if err != nil {

	}

	fileScanner := bufio.NewScanner(currentfile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	currentfile.Close()

	return lines
}

func containsTrigger(line string) bool {

	var i int = 0
	for i < len(triggers) {
		if strings.Contains(line, triggers[i]) {
			return true
		}
		i++
	}

	return false
}
