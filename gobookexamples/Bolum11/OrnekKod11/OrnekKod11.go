package main
import "fmt"
type ogrenci struct {
	ogrID   int
	ogrAdi  string
	ogrSadi string
}
func (o ogrenci) ogrenciYaz() string {
	return fmt.Sprintf("%d numaralı %s %s", o.ogrID, o.ogrAdi, o.ogrSadi)
}
type ogrders struct {
	dersAdi string
	vize    int
	final   int
	ogrenci
}
func (od ogrders) ogrNotYaz() {
	fmt.Println("Öğrenci :", od.ogrenciYaz())
	fmt.Println("Ders Adı :", od.dersAdi)
	fmt.Println("Vize :", od.vize)
	fmt.Println("Final :", od.final)
	fmt.Println("Ortalama :", float32(od.vize)*0.4+float32(od.final)*0.6)
}
type tumdersler struct {
	dersler []ogrders
}
func (d tumdersler) desleriYaz() {
	for _, bilgi := range d.dersler {
		bilgi.ogrNotYaz()
		fmt.Println("=============================")
	}
}
func main() {
	ogr := ogrenci{1, "Ali", "Al"}
	ders1 := ogrders{"Matematik", 72, 80, ogr}
	ders2 := ogrders{"Fizik", 45, 70, ogr}
	ders3 := ogrders{"Kimya", 55, 75, ogr}
	ders4 := ogrders{"Türkçe", 50, 78, ogr}
	ders5 := ogrders{"Sosyal Bilgiler", 68, 78, ogr}
	dersler := tumdersler{[]ogrders{ders1, ders2, ders3, ders4, ders5}}
	dersler.desleriYaz()
}
