package utils

import (
	"log"
	"strconv"
	"strings"
)

func SumFullOverlaps() int {
	s := getScanner("input.txt")
	defer closeOS()

	sum := 0

	for s.Scan() {
		boundA, boundB := getPairBounds(strings.Split(s.Text(), ","))
		if checkFullOverlap(boundA, boundB) {
			sum += 1
		}
	}
	return sum
}

// Takes a string slice with the input format {X-Y, Z-W} and returns a string slice with each range expanded
func getPairBounds(input []string) ([]int, []int) {
	boundA := make([]int, 2)
	boundB := make([]int, 2)

	boundA[0], boundA[1] = getBounds(input[0])

	boundB[0], boundB[1] = getBounds(input[1])

	return boundA, boundB
}

// Returns the lower and upper bound from a string with the format X-Y
func getBounds(s string) (int, int) {
	bounds := strings.Split(s, "-")
	lowerBound, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal(err)
	}

	upperBound, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal(err)
	}

	return lowerBound, upperBound
}

// Returns true if one of the ranges fully overlap the other
func checkFullOverlap(boundA, boundB []int) bool {
	switch {
	case boundA[0] >= boundB[0] && boundA[1] <= boundB[1]:
		return true
	case boundB[0] >= boundA[0] && boundB[1] <= boundA[1]:
		return true
	default:
		return false
	}
}
