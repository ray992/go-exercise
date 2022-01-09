package main

import (
	"fmt"
	"time"
)

func receive(ch1 chan int)  {
	for {
		x := <- ch1
		fmt.Println(x)
	}
}

//单向通道
func print(ch <-chan int)  {
	for i := range ch{
		fmt.Println(i)
	}
}

func write(ch chan<- int)  {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	//CSP 并发模型， 提倡通过通信共享内存，而不是共享内存而实现通信
	var ch1 chan int
	ch1 = make(chan int)
	go receive(ch1)
	fmt.Println(ch1)
	ch1 <- 10
	time.Sleep(time.Second)

	var ch2 chan int
	ch2 = make(chan int, 2)
	ch2 <- 100
	x := <- ch2
	fmt.Println(x)


	ch3 := make(chan int)
	//ch4 := make(chan int)
	go write(ch3)
	go print(ch3)
	time.Sleep(time.Second)
}
