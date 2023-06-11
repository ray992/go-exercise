package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string `json:"myName"` //字段要大写
	Age        int64  `json:"ageInfo"`
	Sex        string `json:"common_sex"`
	Motherland string `json:"motherland"`
}

func main() {

	p := Person{
		Name:       "mike",
		Age:        28,
		Sex:        "male",
		Motherland: "china",
	}
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	var p2 Person
	err = json.Unmarshal(bytes, &p2)
	if err != nil {
		fmt.Errorf("%s", err)
		return
	}
	fmt.Println(p2)
}
