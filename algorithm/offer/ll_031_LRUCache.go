package main

import "fmt"

/** 最近最少使用缓存  **/

type LRUCache struct {
	dataMap map[int]int
	index []int
	capacity int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]int), []int{}, capacity}
}


func (this *LRUCache) Get(key int) int {
	v, ok := this.dataMap[key]
	if ok {
		for i := 0; i < len(this.index); i++ {
			if this.index[i] == key {
				this.index = append(this.index[0:i], this.index[i+1:]...)
				break
			}
		}
		this.index = append(this.index, key)
		return v
	}else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	_, ok := this.dataMap[key]
	if ok {
		for i := 0; i < len(this.index); i++ {
			if this.index[i] == key {
				this.index = append(this.index[0:i], this.index[i+1:]...)
				break
			}
		}
		this.index = append(this.index, key)
		this.dataMap[key] = value // 更新原来的值
	}else {
		size := len(this.dataMap)
		if size >= this.capacity {
			key := this.index[0]
			delete(this.dataMap, key)
			this.index = this.index[1:]
		}
		this.dataMap[key] = value
		this.index = append(this.index, key)
	}
}

func  main()  {
	var a = []int{1,2,3}
	a = append(a[0:1], a[2:]...)
	fmt.Println(a)
	a = a[1:]
	fmt.Println(a)
}