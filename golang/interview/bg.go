package interview

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	var num1 int
	num2 := 1
	f := func() {
		num1, num2 = num2, (num1+num2)%(1e9+7)
	}

	for i := 0; i < n-1; i++ {
		f()
	}

	return num2 % (1e9 + 7)
}
