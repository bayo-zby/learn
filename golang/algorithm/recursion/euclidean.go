package recursion

/*欧几里得算法，辗转相除
funcname:Eucilidean
in:
	@param a int:值1
	@param b int:值2
out:
	@param int:值a与值b的最大公约数
descrip:
	使用递归完成辗转相除
*/
var ch = make(chan int)

func Euclidean(a, b int) int {
	recursionBody(a, b)
	return <-ch
}

func recursionBody(a, b int) {
	if a%b == 0 {
		ch <- b
		return
	}
	go recursionBody(b, a%b)
}
