package slice

import (
	"github.com/LXJ0000/basic-go/lib"
	"github.com/LXJ0000/basic-go/lib/errs"
	"github.com/LXJ0000/basic-go/lib/set"
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

func DeleteAt[T lib.Ordered](nums []T, index int) ([]T, error) {
	if index < 0 || index >= len(nums) {
		return nil, errs.NewErrIndexOutOfRange(len(nums), index)
	}
	return append(nums[:index], nums[index+1:]...), nil
}

func Shrink[T lib.Ordered](nums []T) []T {
	calCapacity := func(c, l int) (int, bool) {
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
	c, l := cap(nums), len(nums)
	n, isChang := calCapacity(c, l)
	if !isChang {
		return nums
	}
	newNums := make([]T, 0, n)
	newNums = append(newNums, nums...)
	return newNums
}

func Unique[T lib.Ordered](nums []T) []T {
	newNums := make([]T, 0, len(nums))
	uniqueByLoop := func(nums []T) {
		for _, num := range nums {
			flag := true
			for _, newNum := range newNums {
				if num == newNum {
					flag = false
				}
			}
			if flag {
				newNums = append(newNums, num)
			}
		}
	}
	uniqueBySet := func(nums []T) {
		s := set.NewMapSet[T](len(nums))
		for _, num := range nums {
			if !s.Exists(num) {
				newNums = append(newNums, num)
			}
			s.Add(num)
		}
	}
	if len(nums) < 1024 {
		uniqueByLoop(nums)
	} else {
		uniqueBySet(nums)
	}
	Shrink(newNums)
	return newNums
}
