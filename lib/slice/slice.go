package slice_

import (
	"github.com/LXJ0000/basic-go/lib"
)

func Max[T lib.Ordered](slice []T) T {
	maxNum := slice[0]
	for _, num := range slice {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func Min[T lib.Ordered](slice []T) T {
	minNum := slice[0]
	for _, num := range slice {
		if num > minNum {
			minNum = num
		}
	}
	return minNum
}

func Sum[T lib.Ordered](slice []T) T {
	var sum T
	for _, k := range slice {
		sum += k
	}
	return sum
}
