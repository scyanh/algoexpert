package recursion

type permutation struct {
}

func NewPermutations() permutation {
	return permutation{}
}

// GetPermutations
/*
<p>
  Write a function that takes in an array of unique integers and returns an
  array of all permutations of those integers in no particular order.
</p>
*/
func (p permutation) GetPermutations(array []int) [][]int {
	permutations := make([][]int, 0)

	p.permutationsHelper(0, array, &permutations)

	return permutations
}

func (p permutation) permutationsHelper(i int, array []int, permutations *[][]int) {
	if i == len(array)-1 {
		newPerm := make([]int, len(array))
		copy(newPerm, array)
		*permutations = append(*permutations, newPerm)
		return
	}

	for j := i; j < len(array); j++ {
		p.swap(array, j, i)
		p.permutationsHelper(i+1, array, permutations)
		p.swap(array, j, i)
	}
}

func (p permutation) swap(array []int, j int, i int) {
	array[j], array[i] = array[i], array[j]
}
