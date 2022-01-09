package main

import (
	"fmt"
	"time"
)

func selectReceive(ch1 chan int, ch2 chan int, ch3 chan int)  {
	for  {
		select {
		case data := <- ch1:
			fmt.Println("ch1 receive data", data)
		case data := <- ch2:
			fmt.Println("ch2 receive data", data)
		case data := <- ch3:
			fmt.Println("ch3 receive data", data)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go selectReceive(ch1, ch2, ch3)
	go func() {
		for m := 1; m <= 10; m++ {
			ch1 <- m
			ch2 <- m*m
			ch3 <- m*m*m
		}
	}()
	time.Sleep(time.Second)
}
