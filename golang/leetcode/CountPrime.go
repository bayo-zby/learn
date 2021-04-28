package leetcode

/*
 * @brief:暴力法计算n在内的素数个数，O(n^2)
 * @param:n,边界值[1,maxint32]
 * @return:素数个数
 */
func CountPrime(n int) int {
	var cot int // 计数器
	// isPrime闭包
	isPrime := func(n int) int {
		// 只要判断到根号n就可以完成判断
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return 0
			}
		}
		return 1
	}
	// 从2开始逐一验证
	for i := 2; i <= n; i++ {
		cot += isPrime(i)
	}
	return cot
}

/*
 * @brief:埃筛法
 * @param:n,判断范围右界，小于等于
 * @ret: int,素数个数
 */
func CountPrimeEratothenes(n int) int {
	// 默认为素数,n+1保证元素和索引能一一对应
	isPrime := make([]bool, n+1)
	var primeCount int
	for i := 2; i <= n; i++ {
		if !(isPrime[i]) {
			// 首次相遇开始计算，不会重复
			primeCount++
			for j := i * i; j < n; j += i {
				isPrime[j] = true
			}
		}
	}
	return primeCount
}
