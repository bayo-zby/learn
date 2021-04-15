package interview_test

import (
	"strings"
	"testing"
)

func strMaker() string {
	var strB strings.Builder
	for i := 0; i < 100000000; i++ {
		strB.WriteByte(byte(i))
	}
	return strB.String()
}

func BenchmarkLenTest(b *testing.B) {
	str := strMaker()
	for i := 0; i < b.N; i++ {
		LenTest(str)
	}
}

func BenchmarkLenTest1(b *testing.B) {
	str := strMaker()
	for i := 0; i < b.N; i++ {
		LenTest1(str)
	}
}

func LenTest(str string) int {
	for i := 0; i < len(str); i++ {

	}
	return len(str)
}

func LenTest1(str string) int {
	l := len(str)
	for i := 0; i < l; i++ {

	}
	return l
}
