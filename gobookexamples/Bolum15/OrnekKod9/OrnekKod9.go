
package main
import "fmt"

func yaz(mesaj string) {
	fmt.Println(mesaj)
}

func panicOrnek(payda int) {
	sonuc := 32 / payda
	fmt.Println("Payda =", sonuc)
}

func main() {
	yaz("Golang Öğreniyorum")
	panicOrnek(0)

}

