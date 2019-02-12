package main
import "fmt"
func sayilariYaz(sonuc chan<- int) {
	for i := 0; i < 10; i++ {
		sonuc <- i
	}
	close(sonuc)
}
func kareAl(sonuc chan<- int, gelen <-chan int) {
	for sayi := range gelen {
		sonuc <- sayi * sayi
	}
	close(sonuc)
}
func ekranaYaz(gelen <-chan int) {
	for sayi := range gelen {
		fmt.Println(sayi)
	}
}
func main() {
	sayilar := make(chan int)
	kareler := make(chan int)

	go sayilariYaz(sayilar)
	go kareAl(kareler, sayilar)
	ekranaYaz(kareler)
}

