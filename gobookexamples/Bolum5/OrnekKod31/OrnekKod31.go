package main

import (
	"fmt"
)

func main() {
	mp1 := make(map[string][]string)

	mp1["key"] = append(mp1["key"], "value1", "value2", "value3")
	fmt.Println(mp1["key"])
}
