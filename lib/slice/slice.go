package slice

import (
	"github.com/LXJ0000/basic-go/lib"
	"github.com/LXJ0000/basic-go/lib/errs"
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
		return nil, errs.NewErrIndexOutOfRange(len(nums), index)
	}
	return append(nums[:index], nums[index+1:]...), nil
}

func Shrink[T any](nums []T) []T {
	c, l := cap(nums), len(nums)
	n, isChang := calCapacity(c, l)
	if !isChang {
		return nums
	}
	newNums := make([]T, 0, n)
	newNums = append(newNums, nums...)
	return newNums
}

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c <= 2048 && (c/l >= 4) {
		return int(float32(c) * 0.5), true
	}
	if c > 2048 && (c/l >= 2) {
		return int(float32(c) * 0.625), true
	}
	return c, false
}
