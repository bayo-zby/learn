package main

import "fmt"

func main() {
<<<<<<< HEAD
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
=======
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		var b int32
		for _, num := range nums {
			// 将每位数字累加
			b += int32(num) >> i & 1
		}
		if n := b % 3; n != 0 {
			ans += n << i
		}
	}
	return int(ans)
>>>>>>> 7df7f931a82d878c1a90e748661fe02ebab4a7be
}
