package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGenerateDocument(t *testing.T) {
	str:="Bste!hetsi ogEAxpelrt x "
	str2:="AlgoExpert is the Best!"

	f := NewGenerateDocument()
	res := f.GenerateDocument(str, str2)

	require.Equal(t, true, res)
}