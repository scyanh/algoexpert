package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRunLengthEncoding(t *testing.T) {
	f := NewRunLengthEncoding()
	res := f.RunLengthEncoding("AAAAAAAAAAAAABBCCCCDD")
	expected := "9A4A2B4C2D"

	require.Equal(t, expected, res)
}
