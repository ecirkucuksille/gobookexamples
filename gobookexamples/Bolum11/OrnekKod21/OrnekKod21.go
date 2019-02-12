package main
import (
	"fmt"
	"reflect"
)
type musteri struct {
	musteriID   int
	musteriAdi  string
	musteriSadi string
}
type urun struct {
	urunID  int
	urunAdi string
	fiyat   float32
}
type siparis struct {
	siparisID int
	musteriID int
	urunID    int
	miktar    int
}
func bilgiVer(arayuz interface{}) {
	fmt.Println("Değişken Tipi =", reflect.TypeOf(arayuz))
	fmt.Println("Değişken Değeri  =", reflect.ValueOf(arayuz))
	fmt.Println("==========================")
}
func main() {
	m := musteri{1, "Ali", "Al"}
	u := urun{1, "Şeker", 4.5}
	s := siparis{1, 1, 1, 2}
	bilgiVer(m)
	bilgiVer(u)
	bilgiVer(s)
}
