package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// inputFile := "inputs/sample.txt"
	// inputFile := "inputs/test.txt"
	inputFile := "inputs/input.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent := string(b)

	res, err := getResult(fileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("The result is", res)
}

func getResult(fileContent string) (int, error) {
	fileContent = strings.Trim(fileContent, " ")
	fileContent = strings.Trim(fileContent, "\n")
	lines := strings.Split(fileContent, ",")

	invalidIdsSum := 0

	for _, line := range lines {
		// remove leading and trailing whitespaces
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}

		// get start and end
		lineSplitted := strings.Split(line, "-")
		startStr := lineSplitted[0]
		endStr := lineSplitted[1]

		// fmt.Println(start)
		// fmt.Println(end)

		start, err := strconv.Atoi(startStr)
		if err != nil {
			return 0, err
		}

		end, err := strconv.Atoi(endStr)
		if err != nil {
			return 0, err
		}

		// theRange := generateRange(start, end)
		// fmt.Println(theRange)

		sumOfInvalidIdsOfRange, err := getSumOfInvalidIdsOfRange(start, end)
		if err != nil {
			return 0, err
		}
		invalidIdsSum += sumOfInvalidIdsOfRange

	}
	return invalidIdsSum, nil
}

func getSumOfInvalidIdsOfRange(start, end int) (int, error) {
	sum := 0
	for num := start; num <= end; num++ {
		numStr := strconv.Itoa(num)
		l := len(numStr)

		// it needs to starts at end with the same digits: must have even length
		if l%2 != 0 {
			continue
		}

		firstHalf := numStr[0 : l/2]
		secondHalf := numStr[l/2 : l]
		if firstHalf == secondHalf {
			sum += num
		}

		// fmt.Println(firstHalf)
		// fmt.Println(secondHalf)

	}
	return sum, nil
}

func generateRange(start, end int) []int {
	var theRange []int
	for num := start; num <= end; num++ {
		theRange = append(theRange, num)
	}
	return theRange
}
