package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int)

func main() {

	rwMu := &sync.RWMutex{}
	num := 1
	go func() {
		for {
			reader(&num, rwMu)
		}
	}()

	go func() {
		for {
			writer(&num, rwMu)
		}
	}()

	<-ch

}

func reader(num *int, rwmu *sync.RWMutex) {
	rwmu.RLock()
	fmt.Println("reader: ", *num)
	if *num == 10 {
		ch <- 0
	}
	time.Sleep(time.Second)
	rwmu.RUnlock()
}

func writer(num *int, rwmu *sync.RWMutex) {

	rwmu.Lock()
	*num++
	fmt.Println("writer: ", *num)
	rwmu.Unlock()
	time.Sleep(time.Second * 2)

}
