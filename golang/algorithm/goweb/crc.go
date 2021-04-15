package goweb

// data为传入数据，r为固定位数
func CRC(data int, r int) int {
	data <<= r
	return Gx(data, 3)
}

// 秦九韶算法
// 二进制转多项式
func Gx(x int, cof int) int {
	exp, x := x%2, x/2
	if x == 0 {
		return exp + cof
	}
	return exp + cof*Gx(x, cof)
}
