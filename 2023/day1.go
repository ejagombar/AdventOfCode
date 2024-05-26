package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func createNumMap() map[string]string {
	m := make(map[string]string)
	m["one"] = "1"
	m["two"] = "2"
	m["six"] = "6"
	m["four"] = "4"
	m["five"] = "5"
	m["nine"] = "9"
	m["three"] = "3"
	m["seven"] = "7"
	m["eight"] = "8"

	return m
}

func checkString(s string, offset int, numMap map[string]string) string {

	lengths := [3]int{3, 4, 5}

	for _, length := range lengths {
		if (offset + length) > len(s) {
			return ""
		}

		str := s[offset : offset+length]
		if num, ok := numMap[str]; ok {
			return num
		}
	}

	return ""
}

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

		numMap := createNumMap()

		left := 0
		stringNum := ""
		for (line[left] < '0' || line[left] > '9') && (stringNum == "") {
			stringNum = checkString(line, left, numMap)
			left++
		}

		var numLeftStr string
		if stringNum != "" {
			numLeftStr = stringNum
		} else {
			numLeftStr = string(line[left])
		}

		right := len(line) - 1
		stringNum = ""
		for (line[right] < '0' || line[right] > '9') && (stringNum == "") {
			stringNum = checkString(line, right, numMap)
			right--
		}

		var numRightStr string
		if stringNum != "" {
			numRightStr = stringNum
		} else {
			numRightStr = string(line[right])
		}

		numStr := numLeftStr + numRightStr

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		total += num
	}

	fmt.Printf("The result is %d \n", total)
}
