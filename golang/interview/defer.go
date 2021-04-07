package interview

// defer在return前
func DoubleScore(a float32) (b float32) {
	defer func() {
		if b >= 100 {
			//将影响返回值
			b = a
		}
	}()

	b = a * 2

	return b
}
