package main
import (
	"fmt"
)
func sayilariYaz(gelen chan int) {
	for i := 1; i <= 5; i++ {
		gelen <- i
	}
	close(gelen)
}
func main() {
	kanal := make(chan int)
	go sayilariYaz(kanal)
	for {
		deger, sonuc := <-kanal
		if sonuc == false {
			break
		}
		fmt.Printf("Okunan Değer = %d, Kanalın Durumu = %t\n", deger, sonuc)
	}
}
