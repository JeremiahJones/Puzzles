package matrix

import (
	"fmt"
	"strings"
)

var (
	coordinates          [][]string
	stringCoordinates    [][]int
	directionCoordinates [][]int
	printableDirections  []string
)

func matrix(givenWord string) {
	startingPosition := "a"
	givenString := startingPosition + givenWord

	//build 2D slice
	row1 := []string{"a", "b", "c", "d", "e"}
	row2 := []string{"f", "g", "h", "i", "j"}
	row3 := []string{"k", "l", "m", "n", "o"}
	row4 := []string{"p", "q", "r", "s", "t"}
	row5 := []string{"u", "v", "w", "x", "y"}
	row6 := []string{"z"}

	coordinates = append(coordinates, row1)
	coordinates = append(coordinates, row2)
	coordinates = append(coordinates, row3)
	coordinates = append(coordinates, row4)
	coordinates = append(coordinates, row5)
	coordinates = append(coordinates, row6)

	//Get coordinate slice from the given string
	for _, v := range givenString {
		letterCoordinates := findLetterCoordinate(string(v), coordinates)
		stringCoordinates = append(stringCoordinates, letterCoordinates)
	}

	//loop through the given coordinate slice and return path
	for i := 0; i < len(stringCoordinates); i++ {
		if i+1 != len(stringCoordinates) {
			results := compareCoordinates(stringCoordinates[i], stringCoordinates[i+1])
			directionCoordinates = append(directionCoordinates, results)
		}
	}

	//print out directions
	for _, val1 := range directionCoordinates {
		directions := []int{val1[0], val1[1]}
		letterDirections := printDirections(directions)
		if letterDirections != "" {
			printableDirections = append(printableDirections, letterDirections)
		}
	}

	fmt.Println(printableDirections)
}

func findLetterCoordinate(givenStringValue string, coordinates [][]string) (indexes []int) {
	for co1, val1 := range coordinates {
		for co2, val2 := range val1 {
			if val2 == givenStringValue {
				indexes = append(indexes, co1, co2)
			}
		}
	}
	return
}

func compareCoordinates(coordinateOne, coordinateTwo []int) (coordinates []int) {
	indexOne := coordinateOne[0] - coordinateTwo[0]
	indexTwo := coordinateOne[1] - coordinateTwo[1]
	coordinates = append(coordinates, indexOne, indexTwo)
	return

}

func printDirections(number []int) string {
	// [0,0] doesnt move, [ (up/-down), (left/-right)]
	var direction string
	var directions []string

	//Down
	if number[0] < 0 {
		direction = "D"
		for i := 0; i < abs(number[0]); i++ {
			directions = append(directions, direction)
		}
	}

	//Up
	if number[0] > 0 {
		direction = "U"
		for i := 0; i < number[0]; i++ {
			directions = append(directions, direction)
		}
	}

	//Left
	if number[1] > 0 {
		direction = "L"
		for i := 0; i < number[1]; i++ {
			directions = append(directions, direction)
		}
	}

	//Right
	if number[1] < 0 {
		direction = "R"
		for i := 0; i < abs(number[1]); i++ {
			directions = append(directions, direction)
		}
	}

	combinedDirections := strings.Join(directions, "")
	return combinedDirections

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
