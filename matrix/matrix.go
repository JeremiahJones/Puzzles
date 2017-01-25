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
		findCoordinateFromString(string(v), Coordinates)
	}
}

func findCoordinateFromString(givenStringValue string, coordinates [][]string) {
	for co1, val1 := range coordinates {
		for co2, val2 := range val1 {
			if val2 == givenStringValue {
				fmt.Println(co1, co2)
			}
		}
	}
}
