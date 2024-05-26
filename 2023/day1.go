package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fileName := argsWithoutProg[0]

	fmt.Print("Opening file ", fileName, "...\n")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		left := 0
		for line[left] < '0' || line[left] > '9' {
			left++
		}

		right := len(line) - 1
		for line[right] < '0' || line[right] > '9' {
			right--
		}

		numStr := string(line[left]) + string(line[right])

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		total += num

	}

	fmt.Printf("The result is %d", total)

}
