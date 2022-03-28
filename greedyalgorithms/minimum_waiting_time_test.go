package greedyalgorithms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinimumWaitingTime(t *testing.T) {
	input := []int{3, 2, 1, 2, 6}
	expectedTime := 17

	m := NewMinimumWaitingTime()
	res := m.MinimumWaitingTime(input)
	require.Equal(t, expectedTime, res)
}
