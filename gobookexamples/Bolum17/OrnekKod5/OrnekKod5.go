package main
import (
	"encoding/json"
	"fmt"
)

//Kisi tipinde yapıyı tanımlar
type Kisi struct {
	Isim    string
	Soyisim string
	Yas     int
}

func main() {
	var kisiler []Kisi
	jsonVeri := `[{"Isim": "Ali", "Soyisim":"Al","Yas": 30}, {"Isim": "Veli", "Soyisim":"Git","Yas": 40}]`
	byts := []byte(jsonVeri)

	hata := json.Unmarshal(byts, &kisiler)
	if hata != nil {
		fmt.Println(hata)
	}

	for sira, veri := range kisiler {
		fmt.Println(sira+1, ". Kişi")
		fmt.Println("==============")
		fmt.Println("Adı :", veri.Isim)
		fmt.Println("Soyadı :", veri.Soyisim)
		fmt.Println("Yaş :", veri.Yas)
	}

}
