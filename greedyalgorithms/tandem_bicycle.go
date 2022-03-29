package greedyalgorithms

import "sort"

type tandemBicycle struct {
}

func NewTandemBicycle() tandemBicycle {
	return tandemBicycle{}
}

func (t tandemBicycle) TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)
	totalSpeed := 0
	if fastest {
		for i := range redShirtSpeeds {
			totalSpeed += t.max(redShirtSpeeds[i], blueShirtSpeeds[len(redShirtSpeeds)-i-1])
		}
	}else{
		for i := range redShirtSpeeds {
			totalSpeed += t.max(redShirtSpeeds[i], blueShirtSpeeds[i])
		}
	}

	return totalSpeed
}

func (tandemBicycle) max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
