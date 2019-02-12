package main

import "fmt"

//Saat tipini  tanımlar
type Saat int

//Dakika tipini tanımlar
type Dakika int

//Kilo tipini tanımlar
type Kilo float64

//Sehir tipini tanımlar
type Sehir string

//Cevap tipini tanımlar
type Cevap bool

func main() {
	saat := Saat(14)
	dakika := Dakika(30)
	kilo := Kilo(45.4)
	sehir := Sehir("Isparta")
	cevap := Cevap(true)

	fmt.Printf("saat  %v:%v\n", saat, dakika)
	fmt.Println("kilo =", kilo)
	fmt.Println("sehir =", sehir)
	fmt.Println("cevap =", cevap)
}
