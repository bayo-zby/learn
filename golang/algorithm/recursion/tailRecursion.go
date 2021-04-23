package recursion

import "fmt"

func test(num int, res int, ch chan int) {
	if num == 1 {
		ch <- res + num
		return
	}
	// 传递计数，计算结果，通道
	// 由于没有进行阻塞，因此创建新的协程后，该协程开辟的内容会在下次GC时被释放
	// 新协程会交由主协程托管
	go test(num-1, res+num, ch)
}

func TestRecursion() {
	ch := make(chan int)
	test(4, 0, ch)
	answer := <-ch
	fmt.Printf("Recursion:%d \n", answer)
}
