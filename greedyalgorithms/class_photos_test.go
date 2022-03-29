package greedyalgorithms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClassPhotos(t *testing.T) {
	redShirtHeights := []int{5, 8, 1, 3, 4}
	blueShirtHeights := []int{6, 9, 2, 4, 5}

	res:=NewClassPhoto().ClassPhotos(redShirtHeights, blueShirtHeights)

	expectedOutput:=true

	require.Equal(t, expectedOutput, res)
}
