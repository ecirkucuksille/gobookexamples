package main

import (
	"fmt"
)

func main() {
	var not = 89
	if not > 90 && not <= 100 {
		fmt.Printf("Harf Notunuz A")
	} else if not > 80 && not <= 90 {
		fmt.Printf("Harf Notunuz B")
	} else if not > 70 && not <= 80 {
		fmt.Printf("Harf Notunuz C")
	} else if not > 60 && not <= 70 {
		fmt.Printf("Harf Notunuz D")
	} else {
		fmt.Printf("Harf Notunuz F")
	}
}
