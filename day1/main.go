package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "inputs/sample.txt"
	// inputFile := "inputs/test.txt"
	// inputFile := "inputs/input.txt"
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
		dialRotation := DialRotation{
			direction:                   c,
			nStepsToRotate:              num,
			dialPosition:                res,
			nTimesRotationBecomes0SoFar: nTimesRotationBecomes0,
		}
		res, nTimesRotationBecomes0, err = rotateTheDial(dialRotation)

		fmt.Println("res after:", res)
		fmt.Println("nTimesRotationBecomes0", nTimesRotationBecomes0)

		fmt.Println()
		fmt.Println()
	}

	return nTimesRotationBecomes0, nil
}

type DialRotation struct {
	direction                   byte
	nStepsToRotate              int
	dialPosition                int
	nTimesRotationBecomes0SoFar int
}

func rotateTheDial(d DialRotation) (dialPosition int, nTimesRotationBecomes0SoFar int, err error) {
	switch d.direction {
	case 'L':
		// convert to negative value to properly add numbers later
		d.nStepsToRotate = d.nStepsToRotate * -1

		// part2, how many times we have crossed the 0
		// Do we cross when we move to the left ?
		if d.dialPosition != 0 && d.dialPosition+d.nStepsToRotate < 0 {
			d.nTimesRotationBecomes0SoFar++
		}

		d.dialPosition = (d.dialPosition + d.nStepsToRotate) % 100
		if d.dialPosition < 0 {
			d.dialPosition += 100
		}

	case 'R':
		// part2, how many times we have crossed the 0
		// Do we cross when we move to the right ?
		if d.dialPosition != 0 && d.dialPosition+d.nStepsToRotate > 100 { // > 100 as we already cover the case 0 later
			d.nTimesRotationBecomes0SoFar++
		}

		d.dialPosition = (d.dialPosition + d.nStepsToRotate) % 100
		if d.dialPosition < 0 {
			d.dialPosition += 100
		}
	default:
		return 0, 0, fmt.Errorf("unsupported rotation direction %v", d.direction)
	}

	switch d.dialPosition {
	case 0:
		d.nTimesRotationBecomes0SoFar++
	case 100:
		d.dialPosition = 0
		d.nTimesRotationBecomes0SoFar++
	}

	// how many full circles have we performed while rotating the dial?
	fullCircles := d.nStepsToRotate / 100
	fmt.Println("temp", fullCircles)
	// always absolute number
	if fullCircles < 0 {
		fullCircles = fullCircles * -1
	}

	if fullCircles > 0 {
		// we don't want to multiply by 0 and loose the count
		d.nTimesRotationBecomes0SoFar += fullCircles

		d.nTimesRotationBecomes0SoFar--
	}
	return d.dialPosition, d.nTimesRotationBecomes0SoFar, nil
}

// 7193 is too high
// 6295 is too low --> last one
