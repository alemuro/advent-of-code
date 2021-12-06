package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getSum(iterator int, path *[]string) int {
	sum := 0

	for i := iterator; i < len(*path) && i < iterator+3; i++ {
		parsed, _ := strconv.Atoi((*path)[i])
		sum += parsed
	}
	return sum
}

func main() {
	lines, _ := readLines("input.txt")

	lastnum := getSum(0, &lines)
	increases := 0

	for i, _ := range lines {
		num := getSum(i, &lines)

		fmt.Printf("%d", num)

		if num > lastnum {
			increases++
			fmt.Printf(" (increase)")
		}

		fmt.Printf("\n")

		lastnum = num
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("%d\n", increases)
	fmt.Printf("--------------------\n")

}
