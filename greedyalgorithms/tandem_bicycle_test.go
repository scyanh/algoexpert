package greedyalgorithms

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTandemBicycle(t *testing.T) {
	redShirtSpeeds := []int{5, 5, 3, 9, 2}
	blueShirtSpeeds := []int{3, 6, 7, 2, 1}
	fastest:=true
	res:=NewTandemBicycle().TandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest)

	expectedSpeed:=32

	require.Equal(t, expectedSpeed, res)
}
