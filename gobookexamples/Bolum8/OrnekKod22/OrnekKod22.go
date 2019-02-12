package main
import "fmt"
func degistir(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
func main() {
	a, b := 3, 5
	fmt.Println("Fonksiyon çağrılmadan önceki değerler")
	fmt.Printf("a = %d, b = %d\n", a, b)
	degistir(&a, &b)
	fmt.Println("Fonksiyon çağrıldıktan sonraki değerler")
	fmt.Printf("a = %d, b = %d\n", a, b)
}
