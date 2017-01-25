package staircase

import (
	"strings"
	"fmt"
)

func printStairCase(number int) (tempSlice []string) {

	for i := 1; i <= number; i++ {

		tempSlice = append(tempSlice, "#")
		stairCase := strings.Join(tempSlice, "")
		fmt.Println(stairCase)
	}
	return tempSlice
}
