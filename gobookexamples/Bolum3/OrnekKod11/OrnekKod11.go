package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b int64 = int64(a)
	fmt.Println(int64(a) + b)
}
