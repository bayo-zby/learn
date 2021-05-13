package leetcode

/**
 * KMP算法
 * 左侧查找右侧语句
 */
func Strstr(source, target string) int {
	if source == "" || target == "" {
		return -1
	}

	ls, lt := len(source), len(target)

	for i := 0; i < ls-lt+1; i++ {
		if source[i:i+lt] == target {
			return i
		}
	}

	return -1
}
