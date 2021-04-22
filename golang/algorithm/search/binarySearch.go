package search
// @program:learn/algorithm/search
// @author:bayo
// @create:2021-04-22 14:50


/*
	该包内函数均为各类查询算法函数，同一文件下为相同算法的不同写法【递归、循环】
*/

/* 二分法查询有序数组
funcname: BinarySearch
input:
	@param:arrInt []int,被查询的数组,默认为从小到大排列
	@param:value int,查找值
out: 
	@param:int ,查找值的索引，未找到则返回-1
description: 利用二分法在有序切片中找到对应值,返回索引值

复杂度为O(logn)
*/
func BinarySearch(arrInt []int, value int) int {
	// 左右指针
	for left, right := 0, len(arrInt)-1; left < right; {
		// 中间量指针
		medianIndex := (left + right) / 2
		switch {
		// 查询值小于中间值，左界移动
		case arrInt[medianIndex] < value:
			left = medianIndex
		// 查询值大于中间值，右界移动个
		case arrInt[medianIndex] > value:
			right = medianIndex
		// 等于则找到对应索引值，直接返回结果
		default:
			return medianIndex
		}
	}
	return -1
}
