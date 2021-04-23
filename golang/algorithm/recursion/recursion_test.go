package recursion_test

import "testing"

func Testeuclidean(t *testing.T) {
	res := recursion.Euclidean(4, 2)
	if res != 2 {
		fmt.Printf("func Euclidean:\nexcept:2,actuall:%d", res)
	}
}
