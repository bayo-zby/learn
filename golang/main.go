package main

import (
	"fmt"
	"math"
)

func main() {
}


func maxSubArray(nums []int) int{
	for i := range(1, len(nums)):
		nums[i] += max(nums[i - 1], 0)
	return max(nums)
}

	