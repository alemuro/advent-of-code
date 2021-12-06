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

func main() {
	lines, _ := readLines("input.txt")

	lastnum, _ := strconv.Atoi(lines[0])
	increases := 0

	for _, e := range lines {
		num, _ := strconv.Atoi(e)

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
