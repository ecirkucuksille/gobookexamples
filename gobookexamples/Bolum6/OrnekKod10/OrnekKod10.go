package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Golang Programlama Go"
	index_1 := strings.Index(s, "Go")
	index_2 := strings.Index(s, "gram")
	index_3 := strings.Index(s, "kod")
	index_4 := strings.LastIndex(s, "Go")
	fmt.Println("index_1 ", index_1)
	fmt.Println("index_2 ", index_2)
	fmt.Println("index_3 ", index_3)
	fmt.Println("index_4 ", index_4)
}
