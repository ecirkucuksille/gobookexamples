package main

import "fmt"

func main() {
	dizi1 := [...]float64{10.2, 54.6, 96.5}
	x := Topla(dizi1)
	fmt.Printf("Dizi elemanlarının toplamı = : %f", x)
}
func Topla(a [3]float64) (toplam float64) {
	for _, v := range a {
		toplam += v
	}
	return
}
