package main

import "fmt"
type Kisi struct {
	ad, soyad string
	yas       int
}
func main() {
	var yapi1 Kisi
	fmt.Println("yapi1 = ", yapi1)

	yapi2 := Kisi{
		ad:  "Ali",
		yas: 25}
	fmt.Println("yapi2 = ", yapi2)
}
