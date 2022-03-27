package dynamicprogramming

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCaseV1(t *testing.T) {
	input := []int{75, 105, 120, 75, 90, 135}
	expected := 330

	s:=NewMaxSubsetSumNoAdjacentService()
	actual := s.MaxSubsetSumNoAdjacent(input)

	require.Equal(t, expected, actual)
}
