package main

import (
	"fmt"
	"runtime"

	"github.com/MrBogomips/stringutil"
)

func main() {
	os := runtime.GOOS
	fmt.Println(os)
	fmt.Println(stringutil.Reverse("Hello World"))
}
