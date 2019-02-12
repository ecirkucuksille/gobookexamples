package main

import (
	"fmt"
)

func main() {
	a := 18
	asalmi := true
	for i := 2; i < a; i++ {
		if a%i == 0 {
			asalmi = false
			break
		}
	}
	if asalmi {
		fmt.Printf("%d , sayısı asaldır.\n", a)
	} else {
		fmt.Printf("%d , sayısı asal değildir.\n", a)
	}
}
