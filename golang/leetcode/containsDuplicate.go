package leetcode

func ContainsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = false
	}
	return true
}
