package main
import (
	"fmt"
)
type myFloat float32
func (a myFloat) bol(b myFloat) myFloat {
	return a / b
}
func main() {
	a := myFloat(5)
	b := myFloat(2)
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a.bol(b))
}
