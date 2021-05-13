package main

import "fmt"

func main() {
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
}
