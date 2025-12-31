package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// inputFile := "inputs/input.txt"
	inputFile := "inputs/sample.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent := string(b)

	initialNumber := 50
	res, err := getResult(fileContent, initialNumber)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("The result is", res)
}

func getResult(fileContent string, initialNumber int) (int, error) {
	fileContent = strings.Trim(fileContent, " ")
	lines := strings.Split(fileContent, "\n")

	res := initialNumber
	nTimesRotationBecomes0 := 0

	for _, line := range lines {
		// remove leading and trailing whitespaces
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}

		fmt.Println("res before:", res)
		fmt.Println(line)

		c := line[0]

		numStr := line[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		switch c {
		case 'L':
			num = num * -1

			// part2, how many times we have crossed the 0
			// first check if we cross the 0 when moving to the left
			if res != 0 && res+num < 0 {
				nTimesRotationBecomes0++
			}
			// how many full circles
			nTimesRotationBecomes0 += (num / 100)

			res = (res + num) % 100
			if res < 0 {
				res += 100
			}

		case 'R':
			// part2, how many times we have crossed the 0
			// Do we cross when we move to the right ?
			if res+num > 100 {
				nTimesRotationBecomes0++
			}

			// how many full circles have we performed while rotating the dial?
			temp := num / 100
			if temp > 0 {
				// we don't want to multiply by 0 and loose the count
				nTimesRotationBecomes0 *= temp
			}

			res = (res + num) % 100
			if res < 0 {
				res += 100
			}
		default:
			return 0, fmt.Errorf("unsupported rotation direction %v", c)
		}
		if res == 0 {
			nTimesRotationBecomes0++
		}

		fmt.Println("res after:", res)
		fmt.Println("crossed the 0 so far:", nTimesRotationBecomes0)

		fmt.Println()
		fmt.Println()
	}

	return nTimesRotationBecomes0, nil
}
