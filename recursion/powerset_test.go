package recursion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPowerset(t *testing.T) {
	arr := []int{1, 2, 3}
	res := NewPowerset().Powerset(arr)

	require.Contains(t, res, []int{})
	require.Contains(t, res, []int{1})
	require.Contains(t, res, []int{2})
	require.Contains(t, res, []int{1, 2})
	require.Contains(t, res, []int{3})
	require.Contains(t, res, []int{1, 3})
	require.Contains(t, res, []int{2, 3})
	require.Contains(t, res, []int{1, 2, 3})
	require.Len(t, res, 8)
}
