package sort

/* 快速排序
funcname : QuickSort
int
	@param sInt []int: 乱序切片
out
	@param 		[]int: 有序切片，从小到达排列
*/
func QuickSort(sInt []int) []int {
	// 基准条件
	if len(sInt) < 2 {
		return sInt
	}
	// 中轴条件
	pivot := sInt[0]
	var less, res, greater []int
	res = append(res, pivot)
	for _, v := range sInt[1:] {
		switch {
		case v < pivot:
			less = append(less, v)
		case v > pivot:
			greater = append(greater, v)
		default:
			res = append(res, v)
		}
	}
	// 分治递归
	less = QuickSort(less)
	greater = QuickSort(greater)

	res = append(less, res...)
	return append(res, greater...)
}
