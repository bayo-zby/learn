package recursion

/*欧几里得算法，辗转相除
funcname:Eucilidean
in:
	@param a int:值1
	@param b int:值2
out:
	@param int:值a与值b的最大公约数
descrip:
	使用协程递归完成辗转相除
*/
var ch = make(chan int)

func Euclidean(a, b int) int {
	go recursionBody(a, b)
	answer := <-ch
	return answer
}

func recursionBody(a, b int) {
	if a%b == 0 {
		ch <- b
		return
	}
	go recursionBody(b, a%b)
}

/*欧几里得算法，辗转相除
funcname:Eucilidean
in:
	@param a int:值1
	@param b int:值2
out:
	@param int:值a与值b的最大公约数
descrip:
	使用正常递归完成辗转相除
*/
func RecursionEuclidean(a, b int) int {
	div := a % b
	if a%b == 0 {
		return b
	}
	return RecursionEuclidean(b, div)
}
