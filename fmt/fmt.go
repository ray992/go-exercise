package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("fmt print")
	fmt.Println()
	fmt.Printf("my name is %s", "mike")
	fmt.Fprintln(os.Stdout, "input")
	file, err := os.OpenFile("./abc.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	msg := "it is test"
	fmt.Fprintf(file, msg)
}
