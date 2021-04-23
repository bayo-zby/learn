package sort

// @ program:go/src/learn/golang/algorithm/sort
// @ author:bayo
// @ create:2021-04-22 15:00

/*
	该包存放各种排序算法，同一文件下为同种排序算法的不同写法
*/

/* 选择排序
funcname:SelectSort
input:
	@ param :arrInt []int,输入数组
out:
	@ param : []int,有序数组，从小到大
description:
	使用原地置换算法，进行选择排序
	复杂度为O(n2)
*/
func SelectSort(arrInt []int) []int {
	for i := 0; i < len(arrInt); i++ {
		var index = i
		for j := i + 1; j < len(arrInt); j++ {
			// 发现更小值，更换最小值索引
			if arrInt[index] > arrInt[j] {
				index = j
			}
		}

		// 当索引变化，则说明有比当前值更小的值
		if index != i {
			arrInt[i], arrInt[index] = arrInt[index], arrInt[i]
		}
	}
	return arrInt
}
