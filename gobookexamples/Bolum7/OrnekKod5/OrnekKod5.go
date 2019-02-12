package main
import "fmt"
func main() {
	var numberOfBytes int
	var err error

	numberOfBytes, err = fmt.Println("Merhaba")

	fmt.Println("err = ", err)
	fmt.Println("numberOfBytes = ", numberOfBytes)
}