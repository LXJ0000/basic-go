package _slice

import "github.com/LXJ0000/basic-go/lib"

func Max[T lib.Ordered](nums ...T) T {
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	return maxNum
}

func Min[T lib.Ordered](nums ...T) T {
	minNum := nums[0]
	for _, num := range nums {
		minNum = max(minNum, num)
	}
	return minNum
}

func Sum[T lib.Ordered](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
