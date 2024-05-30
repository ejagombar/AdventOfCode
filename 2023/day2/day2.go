package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Colours struct {
	red   uint8
	green uint8
	blue  uint8
}

func AddRound(input string) (singleHand Colours, err error) {

	colours := strings.Split(input, ",")

	for _, colour := range colours {
		splitCol := strings.Split(colour, " ")

		if len(splitCol) > 1 {
			num, err := strconv.Atoi(splitCol[1])
			if err != nil {
				return singleHand, fmt.Errorf("AddRound failed %w", err)
			}

			switch splitCol[2] {
			case "red":
				singleHand.red = uint8(num)
			case "green":
				singleHand.green = uint8(num)
			case "blue":
				singleHand.blue = uint8(num)
			}
		}
	}

	return singleHand, nil
}

func ExtractData(file *os.File) ([][]Colours, error) {

	gameStore := make([][]Colours, 0)
	gameNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		string := scanner.Text()
		gameStore = append(gameStore, []Colours{})

		games := strings.Split(string, ":")[1]
		rounds := strings.Split(games, ";")

		for _, x := range rounds {

			round, err := AddRound(x)
			if err != nil {
				return gameStore, fmt.Errorf("ExtractData failed. Error: %w", err)
			}
			gameStore[gameNumber] = append(gameStore[gameNumber], round)
		}
		gameNumber++
	}

	return gameStore, nil
}

func main() {

	argsWithoutProg := os.Args[1:]
	fileName := argsWithoutProg[0]

	fmt.Print("Opening file ", fileName, "...\n")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	gameStore, err := ExtractData(file)
	if err != nil {
		log.Fatal(err)
	}

	IDSum := 0

	for gameIdx, game := range gameStore {
		validGame := true
		for _, round := range game {
			if round.red > 12 || round.green > 13 || round.blue > 14 {
				validGame = false
			}
		}

		if validGame {
			IDSum += (gameIdx + 1)
		}
	}

	fmt.Printf("Sum of valid IDs: %d \n", IDSum)
	powerSum := 0

	for _, game := range gameStore {
		var maxRed, maxGreen, maxBlue uint8
		for _, round := range game {
			maxRed = max(maxRed, round.red)
			maxBlue = max(maxBlue, round.blue)
			maxGreen = max(maxGreen, round.green)
		}
		powerSum += int(maxRed) * int(maxBlue) * int(maxGreen)
	}

	fmt.Printf("Sum of powers: %d \n", powerSum)

}
