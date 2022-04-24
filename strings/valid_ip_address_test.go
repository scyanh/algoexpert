package strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewValidIPAddress(t *testing.T) {
	str := "245245245245"
	v := NewValidIPAddress().ValidIPAddresses(str)

	require.Equal(t, []string{"245.245.245.245"}, v)
}
