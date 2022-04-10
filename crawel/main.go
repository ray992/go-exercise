package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(bytes)
	}
	fmt.Printf("%s", bytes)
}
