package matrix

import (
	"fmt"
)

var Coordinates [][]string

func matrix() {

	givenString := "dogs"

	//build 2D slice
	row1 := []string{"a", "b", "c", "d", "e"}
	row2 := []string{"f", "g", "h", "i", "j"}
	row3 := []string{"k", "l", "m", "n", "o"}
	row4 := []string{"p", "q", "r", "s", "t"}
	row5 := []string{"u", "v", "w", "x", "y"}
	row6 := []string{"z"}

	Coordinates = append(Coordinates, row1)
	Coordinates = append(Coordinates, row2)
	Coordinates = append(Coordinates, row3)
	Coordinates = append(Coordinates, row4)
	Coordinates = append(Coordinates, row5)
	Coordinates = append(Coordinates, row6)

	//loop through the given string slice and iterate to find coordinates
	for _, v := range givenString {
		findCoordinateFromString(v, Coordinates)
	}

	fmt.Println(Coordinates[5][0])
}

func findCoordinateFromString(givenStringValue string, coordinates [][]string) [][]string {
	for ind, val := range coordinates {
		if val == givenStringValue {
			fmt.Println(ind)
			return ind
		}
		return ""
	}
	return ""
}
