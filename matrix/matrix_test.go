package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	matrix("dogs")
}

func TestNewCompareCoordinate(t *testing.T) {
	coordinateOne := []int{3, 3}
	coordinateTwo := []int{4, 4}

	results := compareCoordinates(coordinateOne, coordinateTwo)

	assert := assert.New(t)

	assert.Equal(results, []int{-1, -1})
}

func TestNewPrintDirection(t *testing.T) {
	sampleCoordinate := []int{2, 5}

	coordinates := printDirections(sampleCoordinate)

	assert.Equal(t, "UULLLLL", coordinates)

}

func TestNewWord(t *testing.T) {

}
