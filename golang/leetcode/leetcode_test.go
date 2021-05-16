package leetcode_test

import (
	"learn/golang/leetcode"
	"testing"
)

type ans struct {
	src, tar string
	out      int
}

var strtrArr = []ans{
	{"asdasda", "sda", 1},
}

func TestStrstr(t *testing.T) {
	for _, v := range strtrArr {
		if res := leetcode.Strstr(v.src, v.tar); res != v.out {
			t.Errorf("except:%d,actual:%d\n", v.out, res)
		}
	}
}
