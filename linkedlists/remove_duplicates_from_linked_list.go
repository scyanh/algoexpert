package linkedlists

type removeDuplicatesFromLinkedList struct {
}

func NewRemoveDuplicatesFromLinkedList() removeDuplicatesFromLinkedList {
	return removeDuplicatesFromLinkedList{}
}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// RemoveDuplicatesFromLinkedList
// <p>
//  You're given the head of a Singly Linked List whose nodes are in sorted order
//  with respect to their values. Write a function that returns a modified version
//  of the Linked List that doesn't contain any nodes with duplicate values. The
//  Linked List should be modified in place (i.e., you shouldn't create a brand
//  new list), and the modified Linked List should still have its nodes sorted
//  with respect to their values.
//</p>
// <p>
//  Each 'LinkedList' node has an integer 'value' as well as
//  a 'next' node pointing to the next node in the list or to
//  'None' / 'null' if it's the tail of the list.
//</p>
//  input = 1 -> 1 -> 3 -> 4 -> 4 -> 4 -> 5 -> 6 -> 6
// output = 1 -> 3 -> 4 -> 5 -> 6
func (removeDuplicatesFromLinkedList) RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	current := linkedList

	for current != nil {

		for current.Next != nil && current.Next.Value == current.Value {
			current.Next = current.Next.Next
		}

		current = current.Next
	}

	return linkedList
}

func (removeDuplicatesFromLinkedList) addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = linkedList.Next
	}

	for _, val := range values {
		linkedList := &LinkedList{
			Value: val,
		}

		current.Next = linkedList
		current = current.Next
	}

	return linkedList
}

func (removeDuplicatesFromLinkedList) GetValues(linkedList *LinkedList) []int {
	var res []int

	current := linkedList

	for current != nil {
		res = append(res, current.Value)
		current = current.Next
	}

	return res
}
