package recursion

type powerset struct {
}

func NewPowerset() powerset {
	return powerset{}
}

// Powerset
/*
<p>
  Write a function that takes in an array of unique integers and returns its
  powerset.
</p>
*/
func (powerset) Powerset(array []int) [][]int {
	subsets := [][]int{{}}

	for _, el := range array {
		lenght := len(subsets)
		for i := 0; i < lenght; i++ {
			currentSubset := subsets[i]
			newSubset := append([]int{}, currentSubset...)
			newSubset = append(newSubset, el)
			subsets = append(subsets, newSubset)
		}
	}

	return subsets
}
