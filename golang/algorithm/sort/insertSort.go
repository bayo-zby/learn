package sort

/**
 * @brief	插入排序
 * @param	arr,带排序序列
 */
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		current := arr[i] //当前值复制出来
		// 由于从i左侧是有序的
		// 因此preIndex从i左侧一位开始
		for preIndex >= 0 && arr[preIndex] > current {
			// 如果待插入值小于左侧序列，则说明需要插入到中间位置
			// 所以需要把左侧值右移，继续向中间寻找值
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
}
