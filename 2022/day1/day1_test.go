package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCalorieCountForTopNElves(t *testing.T) {
	testCases := []struct {
		desc     string
		count    int
		expected float64
	}{
		{
			desc:     "Part 1: Top elf calorie count",
			count:    1,
			expected: 69501,
		},
		{
			desc:     "Part 2: Top 3 elves' calorie count",
			count:    3,
			expected: 202346,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			calorieCount := getCalorieCountForTopNElves(testCase.count)
			assert.Equal(t, testCase.expected, calorieCount)
		})
	}
}
