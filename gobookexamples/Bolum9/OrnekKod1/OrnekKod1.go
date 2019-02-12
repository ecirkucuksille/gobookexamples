package main
import "fmt"
func main() {
	b := 3
       var a *int = &b
       fmt.Printf("Veri tipi  %T\n", a)
	fmt.Println("Adres değeri  ", a)
	fmt.Println("Adres içerisindeki değer ",b)
	fmt.Println("Adres içerisindeki değer ",*a)
}
