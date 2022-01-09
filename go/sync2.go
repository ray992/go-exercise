package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x1 int64
	wg1 sync.WaitGroup
	lock1 sync.Mutex
	rwLock sync.RWMutex //读写锁
)

func writeData()  {
	rwLock.Lock() //加写锁
	x1 ++
	time.Sleep(time.Second)
	rwLock.Unlock()
	wg1.Done()
}

func readData()  {
	rwLock.RLock() //加读锁
	time.Sleep(time.Second)
	rwLock.RUnlock()
	wg1.Done()//-1
}

func main() {
	for i := 1; i <= 10; i++ {
		wg1.Add(1)//+1
		go  writeData()
	}
	for i := 1; i <= 10; i++ {
		wg1.Add(1)
		go  readData()
	}
	wg1.Wait() //阻塞
	fmt.Println(x1)
}
