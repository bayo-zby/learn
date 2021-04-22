package leetcode

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	var nums [128]int
	for i := 0; i < l; i++ {
		nums[s1[i]-'A']++
		nums[s2[i]-'A']--
	}

	for _, n := range nums {
		if n != 0 {
			return false
		}
	}

	return true
}
