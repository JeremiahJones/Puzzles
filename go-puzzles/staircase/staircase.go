package staircase

import (
	"fmt"
	"strings"
)

func printStairCase(number int) (tempSlice []string) {

	for i := 1; i <= number; i++ {

		tempSlice = append(tempSlice, "#")
		stairCase := strings.Join(tempSlice, "")
		fmt.Println(stairCase)
	}
	return tempSlice
}
