package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	action string
	units  int
}

func readLines(path string) []Instruction {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineparsed := strings.Split(line, " ")
		num, _ := strconv.Atoi(lineparsed[1])

		lines = append(lines, Instruction{
			action: lineparsed[0],
			units:  num,
		})
	}
	return lines
}

func main() {
	lines := readLines("input.txt")

	var horizontal, depth int

	for _, e := range lines {
		if e.action == "forward" {
			horizontal += e.units
		} else if e.action == "down" {
			depth += e.units
		} else if e.action == "up" {
			depth -= e.units
		}

	}

	fmt.Printf("--------------------\n")
	fmt.Printf("%d\n", horizontal*depth)
	fmt.Printf("--------------------\n")

}
