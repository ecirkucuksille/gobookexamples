package main
import "fmt"
func main() {
	func(mesaj string) {
		fmt.Println(mesaj)
	}("Go öğreniyorum")
}