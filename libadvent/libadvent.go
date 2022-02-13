package libadvent

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func SlurpFile(file string) string {
	input, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func TrimNl(line string) string {
	return strings.TrimSuffix(line, "\n")
}

func InputToInts(input string) []int {
	return linesToInts(splitInput(input))
}

func splitInput(input string) []string {
	return strings.Split(input, "\n")
}

func linesToInts(lines []string) []int {
	var depths []int

	for _, depth := range lines {
		d, _ := strconv.Atoi(depth)
		depths = append(depths, d)
	}

	return depths
}
