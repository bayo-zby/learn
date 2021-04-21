package cache

import (
	"errors"
	"learn/golang/algorithm"
)

type FIFOcache struct {
	size     int
	capacity int // 缓存容量和链表相同
	m        map[string]int
	dln      *algorithm.DoubleListNode
}

func NewFIFO(c int) *FIFOcache {
	if c < -1 {
		c = 0xffff
	}
	return &FIFOcache{
		// 附上容量值
		capacity: c,
		dln:      &algorithm.DoubleListNode{Capacity: c},
	}
}

/* the method is belong to FIFOcache.
 * it can reverse key(string) in FIFOcache.m and try to return the compare value .
 * if the key is not exists , it will return 0 and error
 * @ param input key : string
 * @ param return : int , error
 */
func (fifo *FIFOcache) Get(key string) (int, error) {
	if v, ok := fifo.m[key]; !ok {
		return v, nil
	}
	return 0, errors.New("key is not exists")
}

/*
 *
 */
func (fifo *FIFOcache) Put(key string, value int) error {
	if fifo.capacity == fifo.size {
		return errors.New("the FIFOcache is full")
	}

	return nil
}
