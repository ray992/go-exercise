package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bytes)
}
