package main
import "fmt"
type Ogrenci struct {
	ogr_no int
	ogr_adi string
	ogr_soyadi string
}
func main() {  
	ogrenci := Ogrenci {11235, "Ahmet","Demir"}
	pointer := &ogrenci
	fmt.Println((*pointer).ogr_adi)
	fmt.Println(ogrenci.ogr_adi)
}
