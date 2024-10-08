package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var arr []int

func handler(w http.ResponseWriter, r *http.Request) {
	for n := 0; n < 100000; n++ {
		arr = append(arr, n)
	}

	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
