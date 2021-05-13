package leetcode

/* 辅助数组
 * @param:
 		nums []int : 1 <= nums.length <= 2 * 104,-2^31 <= nums[i] <= 2^31 - 1
		k int : [0,10^5]
 * @return:
*/
// 时间复杂度：O(n),线性增长
// 空间复杂度：O(n),额外新增一个数组空间，空间大小取决于
func Rotate(nums []int, k int) {
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[(i+k)%l] = nums[i]
	}
	copy(nums, res)
}

// 重建数组，移动相当于溢出部分的数组拼接到未溢出部分之后
func ReverseLeftWords(s string, n int) string {
	n = n % len(s)
	return s[n:] + s[:n]
}
