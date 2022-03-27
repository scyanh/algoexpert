package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCaseV1(t *testing.T) {
	input := []int{1, 2, 3, 5, 6, 8, 9}
	expected := []int{1, 4, 9, 25, 36, 64, 81}

	s:=NewSortedSquaredArrayService()
	actual := s.SortedSquaredArray(input)

	require.Equal(t, expected, actual)
}

func TestCaseV2(t *testing.T) {
	input := []int{1, 2, 3, 5, 6, 8, 9}
	expected := []int{1, 4, 9, 25, 36, 64, 81}

	s:=NewSortedSquaredArrayService()
	actual := s.SortedSquaredArrayV2(input)

	require.Equal(t, expected, actual)
}
