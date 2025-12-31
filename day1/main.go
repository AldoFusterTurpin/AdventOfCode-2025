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

	newDialPosition := initialNumber
	nTimesRotationBecomes0 := 0

	for _, line := range lines {
		// remove leading and trailing whitespaces
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}

		// fmt.Println("res before:", newDialPosition)
		// fmt.Println(line)

		c := line[0]

		numStr := line[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}
		dialRotation := DialRotation{
			direction:      c,
			nStepsToRotate: num,
			dialPosition:   newDialPosition,
		}
		var nTimesCrossed0InThisRotation int
		newDialPosition, nTimesCrossed0InThisRotation, err = rotateTheDial(dialRotation)
		nTimesRotationBecomes0 += nTimesCrossed0InThisRotation

		// fmt.Println("res after:", newDialPosition)

		// fmt.Println()
		// fmt.Println()
	}

	return nTimesRotationBecomes0, nil
}

type DialRotation struct {
	direction      byte
	nStepsToRotate int
	dialPosition   int
}

func rotateTheDial(d DialRotation) (newDialPosition int, nTimesCrossed0InThisRotation int, err error) {
	newDialPosition = d.dialPosition
	nTimesCrossed0InThisRotation = 0

	// offsetWithSign will be used to calculate the final position we will land on.
	// It will be added (+) to the initial position, so we will convert it to negative for the 'L' case
	// casue we want to always do n + (offsetWithSign)
	offsetWithSign := d.nStepsToRotate

	// direction will be used for part 2 to move to left or right
	direction := 0
	switch d.direction {
	case 'L':
		direction = -1
		// convert to negative value to properly add numbers later (avoiding - -x becomes +)
		offsetWithSign = offsetWithSign * -1
	case 'R':
		direction = 1
	default:
		return 0, 0, fmt.Errorf("unsupported rotation direction %v", d.direction)
	}

	// Part 1: compute final position
	newDialPosition = (newDialPosition + offsetWithSign) % 100

	// we want the modulo to always be positive
	if newDialPosition < 0 {
		newDialPosition += 100
	}

	// part 2: how many times we cross the 0
	// Easiest thing to do is to simply move in the direction and convert values as we move.
	// We could use some math triks to simplify the problem but this is O(n) and simple to reason about.
	current := d.dialPosition

	// we simply move d.nStepsToRotate times in the given direction calculated before
	for i := 0; i < d.nStepsToRotate; i++ {
		current += direction
		switch current {
		case 100:
			current = 0
		case -1:
			current = 99
		}
		if current == 0 {
			nTimesCrossed0InThisRotation++
		}
	}

	return newDialPosition, nTimesCrossed0InThisRotation, nil
}
