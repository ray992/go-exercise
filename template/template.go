package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//
	template, err := template.ParseFiles("./template/hello.tmpl")
	if err != nil {
		fmt.Println("create failed", err)
	}
	template.Execute(w, "mike")
}

func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		fmt.Println("http server failed", err)
		return
	}
}
