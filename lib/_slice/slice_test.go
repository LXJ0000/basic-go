package _slice

import (
	"testing"

	"github.com/LXJ0000/basic-go/lib"
)

func TestMaxMinSum(t *testing.T) {
	testCases := []struct {
		nums     []lib.Ordered
		expected lib.Ordered
	}{
		{[]lib.Ordered{1, 2, 3}, 3},             // Max test case
		{[]lib.Ordered{-1, -2, -3}, -1},         // Max test case with negative numbers
		{[]lib.Ordered{5, 5, 5}, 5},             // Max test case with equal numbers
		{[]lib.Ordered{10, 20, 30}, 10},         // Min test case
		{[]lib.Ordered{-10, -20, -30}, -30},     // Min test case with negative numbers
		{[]lib.Ordered{5, 5, 5}, 5},             // Min test case with equal numbers
		{[]lib.Ordered{1, 2, 3}, 6},             // Sum test case
		{[]lib.Ordered{-1, -2, -3}, -6},         // Sum test case with negative numbers
		{[]lib.Ordered{10, 20, 30}, 60},         // Sum test case with positive numbers
		{[]lib.Ordered{1.1, 2.2, 3.3}, 3.3},     // Max test case with float numbers
		{[]lib.Ordered{-1.1, -2.2, -3.3}, -1.1}, // Max test case with negative float numbers
		{[]lib.Ordered{1.1, 2.2, 3.3}, 1.1},     // Min test case with float numbers
		{[]lib.Ordered{-1.1, -2.2, -3.3}, -3.3}, // Min test case with negative float numbers
		{[]lib.Ordered{1.1, 2.2, 3.3}, 6.6},     // Sum test case with float numbers
		{[]lib.Ordered{-1.1, -2.2, -3.3}, -6.6}, // Sum test case with negative float numbers
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			// Test Max function
			maxResult := Max(tc.nums...)
			if maxResult != tc.expected {
				t.Errorf("Max result for %v is incorrect, got: %v, want: %v", tc.nums, maxResult, tc.expected)
			}

			// Test Min function
			minResult := Min(tc.nums...)
			if minResult != tc.expected {
				t.Errorf("Min result for %v is incorrect, got: %v, want: %v", tc.nums, minResult, tc.expected)
			}

			// Test Sum function
			sumResult := Sum(tc.nums...)
			if sumResult != tc.expected {
				t.Errorf("Sum result for %v is incorrect, got: %v, want: %v", tc.nums, sumResult, tc.expected)
			}
		})
	}
}
