package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseWordsInString(t *testing.T) {
	r:=NewReverseWordsInString().ReverseWordsInString("a b c")
	require.Equal(t, "c b a", r)
}
