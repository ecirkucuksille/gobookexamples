package ogrenci
import "fmt"
//Ogrenci yapısını tanımlar
type ogrenci struct {
	ogrID   int
	ogrAdi  string
	ogrSadi string
	vize    int
	final   int
}
//New fonksiyonu kurucu fonksiyon gibi çalışır
func New(ogrID int, ogrAdi string, ogrSadi string, vize int, final int) ogrenci {
	o := ogrenci{ogrID, ogrAdi, ogrSadi, vize, final}
	return o
}
//OrtalamaHesapla ortalamayı hesaplar
func (o ogrenci) OrtalamaHesapla() {
	ortalama := float32(o.vize)*0.4 + float32(o.final)*0.6
	fmt.Printf("%d numaralı %s %s'nin ortalaması = %.2f'dir\n", o.ogrID, o.ogrAdi, o.ogrSadi, ortalama)
}
