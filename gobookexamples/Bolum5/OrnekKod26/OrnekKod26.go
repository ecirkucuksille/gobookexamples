package main

import "fmt"

func Topla(a []float64) float64 {
	s := 0.0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
func main() {
	var dizi1 = [3]float64{10.2, 54.6, 96.5}
	fmt.Println("Dizinin elemanları toplamı = ", Topla(dizi1[:]))
}
