package main
import (
	"fmt"
)
func kontrol(i interface{}) {
	deger := i.(int)
	fmt.Println("Gelen Değer =", deger)
}
func main() {
	sayi := 3.14
	kontrol(sayi)
}