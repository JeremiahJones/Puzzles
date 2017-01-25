package staircase

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewStaircase(t *testing.T) {
	numberOfSteps := 4
	t.Log("Building staircase with &s steps", numberOfSteps)
 	stairCase := printStairCase(numberOfSteps)

	assert := assert.New(t)
	assert.Equal([]string{"#", "#", "#", "#"},stairCase)

}
