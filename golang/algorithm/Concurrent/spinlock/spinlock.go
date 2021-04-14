package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLock uint32

func (sl *SpinLock) Lock() {
	// 比较新旧值，替换，当sl的值不为0时,延缓当前协程执行
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *SpinLock) Unlock() {
	// 将sl的值修改为0,该操作为原子操作
	atomic.StoreUint32((*uint32)(sl), 0)
}

// 返回接口
func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}

func main() {
}

func reader()
