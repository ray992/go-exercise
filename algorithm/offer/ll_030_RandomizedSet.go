package main

import (
	"fmt"
	"math/rand"
)

/** 插入、删除和随机访问都是 O(1) 的容器 **/

type RandomizedSet struct {
	dataMap   map[int]int
	indexMap  map[int]int
	index     []int
	last int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int, 0), make(map[int]int, 0), []int{}, 0}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dataMap[val]
	if ok {
		return false
	}
	this.index = append(this.index, this.last)
	this.dataMap[val] = this.last
	this.indexMap[this.last] = val
	this.last++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	last, ok := this.dataMap[val]
	if !ok {
		return false
	}
	delete(this.dataMap, val)
	delete(this.indexMap, last)
	l, r := 0, len(this.index) - 1
	for l <= r {
		mid := (r - l)/2 + l
		if this.index[mid] == last {
			this.index = append(this.index[:mid], this.index[mid+1:]...)
			break
		}else if this.index[mid] > last {
			r--
		}else {
			l++
		}
	}
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.index))
	v, _ := this.indexMap[this.index[random]]
	return v
}

func main(){
	fmt.Println(rand.Intn(10))
}