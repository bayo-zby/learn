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

// 原地置换
func Rotate2(nums []int, k int) {

}
