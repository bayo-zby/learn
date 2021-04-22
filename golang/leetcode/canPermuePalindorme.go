package leetcode

func CanPermutePalindrome(s string) bool {
	strMap := make(map[rune]int)
	for _, v := range s {
		strMap[v]++
	}
	odd := 0
	for _, v := range strMap {
		if v%2 != 0 {
			odd++
		}
	}
	return odd <= 1
}
