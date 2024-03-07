package _slice

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// go get "github.com/stretchr/testify/require"

func TestMax(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -1,
		},
		{
			name:     "mixed numbers",
			nums:     []int{-10, 0, 10, -5, 5},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Max(tc.nums...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -5,
		},
		{
			name:     "mixed numbers",
			nums:     []int{-10, 0, 10, -5, 5},
			expected: -10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Min(tc.nums...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -15,
		},
		{
			name:     "mixed numbers",
			nums:     []int{-10, 0, 10, -5, 5},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.nums...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestDeleteAt(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		index    int
		expected []int
		err      error
	}{
		{
			name:     "positive",
			nums:     []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
			err:      nil,
		},
		{
			name:     "negative",
			nums:     []int{-5, -4, -3, -2, -1},
			index:    -1,
			expected: nil,
			err:      errors.New("index out of range"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DeleteAt(tc.nums, tc.index)
			require.Equal(t, tc.expected, result)
			require.Equal(t, tc.err, err)
		})
	}
}
