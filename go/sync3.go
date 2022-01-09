package main

import (
	"fmt"
	"sync"
)

type mate struct {
	name string
	age int
}

var instance *mate
var once sync.Once

func initMate()  *mate {
	once.Do(
		func() {
			instance = &mate{"mike", 28}
		})
	return instance
}

func main() {
	curMate := initMate()
	fmt.Println(*curMate)
}
