package main
import (
	"fmt"
	"sync"
	"time"
)
func main() {

	rutinSayisi := 3
	var wg sync.WaitGroup

	wg.Add(rutinSayisi)

	for i := 1; i <= rutinSayisi; i++ {
		go esZamanliIslemler(i, &wg)
	}

	wg.Wait()
	fmt.Println("Bitti")
}
func esZamanliIslemler(sure int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(sure) * time.Second)
	fmt.Println(sure, "saniye beklendi")
}

