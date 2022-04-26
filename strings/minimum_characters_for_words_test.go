package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinimumCharactersForWords(t *testing.T) {
	str := []string{"this", "that", "did", "deed", "them!", "a"}
	res := NewMinimumCharactersForWords().MinimumCharactersForWords(str)
	expected := []string{"m", "s", "t", "t", "i", "e", "e", "!", "h", "a", "d", "d"}

	require.ElementsMatch(t, expected, res)

}
