package main

import (
	"fmt"
)

func main() {
	sayi := 64
	for i := 2; i <= sayi; i++ {
		if sayi%i == 0 {
			fmt.Printf("%d , %d 'ün bölenidir.\n", i, sayi)
		}
	}
}
