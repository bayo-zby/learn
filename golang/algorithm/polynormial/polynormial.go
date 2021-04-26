package polynormial

// 多项式包

/* 秦九韶算法
二进制转多项式
@ param x int,指数项
@ param cof int,底数
*/
func Gx(x int, cof int) int {
	exp, x := x%2, x/2
	if x == 0 {
		return exp + cof
	}
	return exp + cof*Gx(x, cof)
}
