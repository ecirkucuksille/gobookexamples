package main
import (
	"fmt"
	"time"
)
func sayilariYaz() {
	for i := 1; i < 6; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(i)
	}
}
func harfleriYaz() {
	for _, v := range "selam" {
		time.Sleep(350 * time.Millisecond)
		fmt.Print(string(v))
	}
}
func main() {
	go harfleriYaz()
	go sayilariYaz()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Program Sonu")
}
