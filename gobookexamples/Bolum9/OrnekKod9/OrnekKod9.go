package main
import (
	"fmt"
	"strings"
)
func buyukYaz (val *string){
	(*val) =strings.ToUpper(*val)
}
func main() {
	a := "Hello World"
             fmt.Println("fonksiyondan önce  :",a)
             b := &a
            buyukYaz(b)
            fmt.Println("fonksiyondan sonra :",a)
}
