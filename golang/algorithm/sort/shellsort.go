package sort

/**
 * @date:	2021-05-11
 * @author:	bayo
 * @brief:	插入排序，原地置换，时间复杂度O(n log n)
 * @param:	arr,待排序序列
 */
func ShellSort(arr []int) {
	//希尔排序的基本思想是：
	//1. 先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序.
	//2.待整个序列中的记录"基本有序"时，再对全体记录进行依次直接插入排序。
	var gap int
	// 动态定义间隔序列
	gap = len(arr)/3 + 1
	for gap > 0 {
		for i := gap; i < len(arr); i++ {
			j := i - gap
			temp := arr[i]
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
}
