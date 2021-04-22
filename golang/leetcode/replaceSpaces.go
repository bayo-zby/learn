package leetcode

import "strings"

// Builder缓冲区
func ReplaceSpaces(S string, length int) string {
	var b strings.Builder
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			b.WriteString("%20")
			continue
		}
		b.WriteByte(S[i])
	}
	return b.String()
}

// 创建数组
func ReplaceSpaces1(S string, length int) string {
	var s = make([]byte, 0, length)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			s = append(s, '%', '2', '0')
			continue
		}
		s = append(s, S[i])
	}
	return string(s)
}
