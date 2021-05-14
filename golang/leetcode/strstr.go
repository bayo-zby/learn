package leetcode

/**
 * 左侧查找右侧语句
 */
func Strstr(source, target string) int {
	// 临界值考虑
	if source == "" || target == "" {
		return -1
	}

	ls, lt := len(source), len(target)

	// 剩余长度不足以覆盖目标字段则结束
	for i := 0; i < ls-lt+1; i++ {
		var j int
		for ; j < lt; j++ {
			if source[i+j] != target[j] {
				j = -1 // 异常返回标记
				break
			}
		}

		// 判断j值是否正常
		if j < 0 {
			return i
		}
	}

	return -1
}
