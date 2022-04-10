package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "my email is yq15194pp@gmail.com \n" +
		"my email is yq15194pp@163.com"
	mustRe := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`) //+至少1个
	match := mustRe.FindAllString(text, -1)
	fmt.Println(match)
}
