package main

import "fmt"

func main() {
	a := []int{1, 1, 4, 5, 6, 2, 2, 2, 2, 2, 2}
	fmt.Println(majorityElement(a))
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var slice []int
	slice = append(slice, nums[0])
	for i := 1; i < len(nums); i++ {
		if len(slice) == 0 {
			slice = append(slice, nums[i])
			continue
		}
		if slice[0] == nums[i] {
			slice = append(slice, nums[i])
		} else {
			slice = slice[1:]
		}
	}
	return slice[0]
}
