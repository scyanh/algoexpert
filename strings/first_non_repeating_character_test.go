package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFirstNonRepeatingCharacter(t *testing.T) {
	str := "abcdcaf"

	res := NewFirstNonRepeatingCharacter().FirstNonRepeatingCharacter(str)
	require.Equal(t, 1, res)
}
