package linkedlists

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRemoveDuplicatesFromLinkedList(t *testing.T) {
	s:=NewRemoveDuplicatesFromLinkedList()
	input := s.addMany(&LinkedList{Value: 1}, []int{1, 3, 4, 4, 4, 5, 6, 6})
	expected := s.addMany(&LinkedList{Value: 1}, []int{3, 4, 5, 6})
	actual := s.RemoveDuplicatesFromLinkedList(input)
	require.Equal(t, s.GetValues(expected), s.GetValues(actual))
}
