package main
import "fmt"
func main() {
	y := 255
	s := "Go Programlama"
	var a *int = &y
	var b *string =&s
	fmt.Println("y=",*a)
	fmt.Println("s=",*b)
	y +=10
	*a *=2
	*b="Pointer"
	fmt.Println("y=",y)
	fmt.Println("s=",*b)
}
