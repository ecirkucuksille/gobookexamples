package main
import "fmt"
func ekranaYaz(sayilar ...int) {
	defer fmt.Println("Yazma İşlemi Sona Erdi")
	for _, sayi := range sayilar {
		defer fmt.Println(sayi)
	}
	fmt.Println("Yazma İşlemi Başladı")
}
func main() {
	ekranaYaz(1, 2, 3, 4, 5)
}
