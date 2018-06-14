package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	os := runtime.GOOS
	fmt.Println(os)
	fmt.Println("Hello World")
}
