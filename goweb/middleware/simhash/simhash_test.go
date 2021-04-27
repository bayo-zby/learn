package simhash_test

import (
	"learn/goweb/middleware/simhash"
	"testing"
)

func TestWordWeight(t *testing.T) {
	words := simhash.WordWeight("我来到北京清华大学")
	if len(words) != 0 {
		t.Errorf("Except : %v,Actual : %v\n", words, words)
	}
}
