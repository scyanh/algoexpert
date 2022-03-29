package greedyalgorithms

import "sort"

type classPhoto struct {
}

func NewClassPhoto() classPhoto {
	return classPhoto{}
}

func (classPhoto) ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Slice(redShirtHeights, func(i, j int) bool {
		return redShirtHeights[i] > redShirtHeights[j]
	})

	sort.Slice(blueShirtHeights, func(i, j int) bool {
		return blueShirtHeights[i] > blueShirtHeights[j]
	})

	redShirtInBack := false
	if redShirtHeights[0] > blueShirtHeights[0] {
		redShirtInBack = true
	}

	for i := range redShirtHeights{
		if redShirtInBack {
			if redShirtHeights[i] <= blueShirtHeights[i]{
				return false
			}
		}else{
			if blueShirtHeights[i] <= redShirtHeights[i]{
				return false
			}
		}
	}

	return true
}
