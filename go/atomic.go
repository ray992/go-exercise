package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter interface {
	Inc()
	Load() int64
}

type atomicCounter struct {
	counter int64
}

func (a *atomicCounter) Inc()  {
	atomic.AddInt64(&a.counter, 1)
}

func (a *atomicCounter) Load() int64  {
	return atomic.LoadInt64(&a.counter)
}

func test(c counter)  {
	var wg2 sync.WaitGroup
	for i := 1; i <= 100; i++{
		wg2.Add(1)
		go func() {
			c.Inc()
			wg2.Done()
		}()
	}
	wg2.Wait()
}

//原子操作
func main() {
	ac := atomicCounter{100}
	test(&ac)
	fmt.Println(ac.counter)
}
