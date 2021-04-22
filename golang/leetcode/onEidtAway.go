package leetcode

func OneEditAway(first string, second string) bool {
	if first == second {
		return true
	}

	l1 := len(first)
	l2 := len(second)
	// 绝对值小于等于1
	if l1-l2 > 1 || l1-l2 < -1 {
		return false
	}

	var strCount [128]int
	for i := 0; i < l1 || i < l2; i++ {
		if i < l1 {
			strCount[first[i]-'A']++
		}
		if i < l2 {
			strCount[second[i]-'A']--
		}
	}

	c := 0
	for _, v := range strCount {
		if v != 0 {
			c++
		}
	}

	return c < 3
}
