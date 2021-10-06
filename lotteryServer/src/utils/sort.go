package utils

import "math"

func MinEleInds(arr *[]int) *[]int {
	minVal := math.MaxInt32
	inds := []int{}

	for k, v := range *arr {
		if minVal == v {
			inds = append(inds, k)
		} else if minVal > v {
			minVal = v
			inds = []int{k}
		}
	}
	return &inds
}