package recursion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPermutation(t *testing.T) {
	arr:=[]int{1, 2, 3}

	// expected:=[][]int{{1, 2, 3},{1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	res := NewPermutations().GetPermutations(arr)
	require.Contains(t, res, []int{1, 2, 3})
	require.Contains(t, res, []int{1, 3, 2})
	require.Contains(t, res, []int{2, 1, 3})
	require.Contains(t, res, []int{2, 3, 1})
	require.Contains(t, res, []int{3, 1, 2})
	require.Contains(t, res, []int{3, 2, 1})
	require.Len(t, res, 6)
}
