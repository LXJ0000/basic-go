package _slice

import (
	"errors"

	"github.com/LXJ0000/basic-go/lib"
)

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
		minNum = min(minNum, num)
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

func DeleteAt[T any](nums []T, index int) ([]T, error) {
	if index < 0 || index >= len(nums) {
		return nil, errors.New("index out of range")
	}
	return append(nums[:index], nums[index+1:]...), nil
}
