package cache

import "learn/golang/algorithm/dataStruct"

type LRUcache struct {
	Capacity int
	List     *dataStruct.DoubleListNode
}

/*
 */
func (lru *LRUcache) Get(key string) (v int) {
	// 获取对应节点
	// 转移到头部
	// 对应位置节点对接
	// 返回值
	return
}

/*
 */
func (lru *LRUcache) Put(key string, value int) {
	// 判断是否存在缓存中
	// 存在则提出旧节点
	// 不存在则
}
