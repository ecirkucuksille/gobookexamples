package main
import (
	"fmt"
	"time"
)
func main() {
	kanal := make(chan bool, 1)

	select {
	case deger := <-kanal:
		fmt.Println(deger)
	case <-time.After(2 * time.Second):
		fmt.Println("2 sn içerisinde veri gelmedi")
	}
}
