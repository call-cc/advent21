package libadvent

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func SlurpFile() string {
	input, err := ioutil.ReadFile(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func TrimNl(line string) string {
	return strings.TrimSuffix(line, "\n")
}

func Convert(lines []string) []int {
	var depths []int

	for _, depth := range lines {
		d, _ := strconv.Atoi(depth)
		depths = append(depths, d)
	}

	return depths
}
