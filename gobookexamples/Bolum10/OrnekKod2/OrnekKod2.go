package main
import "fmt"
type Kisi struct {
	ad, soyad string
	yas       int
}
func main() {
	yapi := Kisi{
		ad:    "Ali",
		yas:   25,
		soyad: "Al",
	}
	fmt.Println("Struct 1 = ", yapi)
	yapi2 := Kisi{"Mehmet", "Demir", 36}
	fmt.Println("Struct 2 = ", yapi2)

	var yapi3 Kisi
	yapi3.ad = "Ayşe"
	yapi3.yas = 52
	yapi3.soyad = "Yıldız"
	fmt.Println("Struct 3 = ", yapi3)
}
