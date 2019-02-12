package main

import "fmt"

func main() {
	sayilar := [5]int{5, 10, 15, 20, 25}
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}
	fmt.Print("Sayılar dizisinin elemanları toplamı = ", toplam)
}
