package sort

/**
 * BullbleSort
 * @brief	冒泡排序，时间复杂度O(n^2)，空间复杂度O(1),从小到大排列
 * @param	数组，冒泡使用原地置换，所以不需要返回值
 */
func BubbleSort(arr []int) {
	var cycleNum int // 循环轮次
	for i := 0; i < len(arr)-cycleNum; {
		for j := i + 1; j < len(arr)-cycleNum; j++ {
			// 将较大的数浮到
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		cycleNum++
	}
}
