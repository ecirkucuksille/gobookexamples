
package main
import (
	"fmt"
)
func bolme(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Sıfıra Bölme İşlemi: %.2f / %.2f", a, b)
	}
	return a / b, nil
}
func main() {
	a, b := 3.0, 0.0
	sonuc, hata := bolme(a, b)
	if hata == nil {
		fmt.Println("a / b =", sonuc)
	} else {
		fmt.Println(hata)
	}
}

