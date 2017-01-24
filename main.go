package matrix

import (
	"fmt"
)

type Grid struct {
	Coordinates [][]string `"json:coordinates"`
}

func matrix() {
	grid := Grid{}
	givenString := "dogs"

	//build 2D slice
	row1 := []string{"a", "b", "c", "d", "e"}
	row2 := []string{"f", "g", "h", "i", "j"}
	row3 := []string{"k", "l", "m", "n", "o"}
	row4 := []string{"p", "q", "r", "s", "t"}
	row5 := []string{"u", "v", "w", "x", "y"}
	row6 := []string{"z"}

	grid.Coordinates = append(grid.Coordinates, row1)
	grid.Coordinates = append(grid.Coordinates, row2)
	grid.Coordinates = append(grid.Coordinates, row3)
	grid.Coordinates = append(grid.Coordinates, row4)
	grid.Coordinates = append(grid.Coordinates, row5)
	grid.Coordinates = append(grid.Coordinates, row6)

	//loop through the given string slice and iterate to find coordinates
	for _, v := range givenString {
		findCoordinateFromString(v, grid)
	}

	fmt.Println(grid[5][0])
}

func findCoordinateFromString(givenStringValue string, coordinates [][]string) [][]string {
	for ind, val := range coordinates {
		if val == givenStringValue {
			return ind
		}
		return ""
	}
}
