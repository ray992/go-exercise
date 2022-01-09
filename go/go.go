package main

import (
	"fmt"
	"runtime"
	"time"
)

func hello()  {
	fmt.Println("hello")
}


//goroutine的调度是在用户态下完成的， 不涉及内核态和用户之间的频繁切换，包括内存的分配与释放

func a()  {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b()  {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	go hello()
	fmt.Println("main goroutine")
	time.Sleep(time.Second)

	runtime.GOMAXPROCS(2) //当前并发时占用的CPU逻辑核心数
	go a()
	go b()
	time.Sleep(time.Second)
}
