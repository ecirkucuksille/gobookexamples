package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	var kelime string
	var sayi float32

	okuyucu := bufio.NewReader(os.Stdin)
	fmt.Print("Kelimeyi Giriniz : ")
	kelime, _ = okuyucu.ReadString('\n')
	kelime = strings.Replace(kelime, "\n", "", -1)
	fmt.Print("Sayıyı Giriniz : ")
	fmt.Scan(&sayi)

	fmt.Println("=================================")

	fmt.Println("Kelime = ", kelime)
	fmt.Println("Sayı = ", sayi)
}