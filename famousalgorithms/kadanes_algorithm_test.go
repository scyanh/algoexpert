package famousalgorithms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCaseKadanesAlgorithm(t *testing.T) {
	input := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}

	f := NewKadanesAlgorith()
	res := f.KadanesAlgorithm(input)
	expectedMaxNumber := 19
	require.Equal(t, expectedMaxNumber, res)
}
