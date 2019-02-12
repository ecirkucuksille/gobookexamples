package main
import "fmt"
func faktoriyel(sayi int) int {
	sonuc := 1
	for i := 1; i <= sayi; i++ {
		sonuc *= i
	}
	return sonuc
}
func main() {
	var girilenSayi int
	fmt.Print("Faktöriyeli Alınacak sayıyı Giriniz : ")
	fmt.Scan(&girilenSayi)
	fmt.Printf("%d! = %d\n", girilenSayi, faktoriyel(girilenSayi))
}