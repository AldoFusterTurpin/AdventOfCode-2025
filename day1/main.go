package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "inputs/input.txt"
	// inputFile := "inputs/sample.txt"
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
	// fmt.Println(fileContent)
	lines := strings.Split(fileContent, "\n")

	res := initialNumber
	nTimesRotationBecomes0 := 0

	for _, line := range lines {
		// remove leading and trailing whitespaces
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}

		c := line[0]
		// fmt.Println(string(c))

		numStr := line[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}
		// fmt.Println(num)

		switch c {
		case 'L':
			num = num * -1
			res = (res + num) % 100
			if res < 0 {
				res += 100
			}
		case 'R':
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

		// fmt.Println("res so far:", res)
		// fmt.Println("***************************************")
	}
	return nTimesRotationBecomes0, nil
}
