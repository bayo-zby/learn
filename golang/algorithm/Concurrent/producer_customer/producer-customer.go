package main

import (
	"fmt"
	"time"
)

/*生产者-消费者模型
使用互斥量解决消费者问题
*/
var mu = make(chan struct{})

func main() {
	ch := make(chan struct{}, 5)

	go custome(ch)
	go produce(ch)

	<-mu
	fmt.Println("结束")
}

// 生产者
func produce(ch chan struct{}) {
	for i := 0; i < 20; i++ {
		ch <- struct{}{}
		fmt.Printf("产品入库，当前数量：%d\n", len(ch))
	}

}

// 消费者
func custome(ch chan struct{}) {
	for {
		<-ch
		fmt.Printf("产品出库，当前数量：%d\n", len(ch))
		time.Sleep(time.Second)
	}
}
