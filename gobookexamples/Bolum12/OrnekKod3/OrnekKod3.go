package main
import (
	"flag"
	"fmt"
)
func main() {
	kelimePtr := flag.String("kelime", "gelmedi", "string bir değişken")
	tekrarPtr := flag.Int("tekrar", 1, "int bir değişken")
	durumPtr := flag.Bool("durum", false, "bool bir değişken")
	var sonYazi string
	flag.StringVar(&sonYazi, "sonYazi", "son yazı yok", "bir string değişken")
	flag.Parse()

	fmt.Println("kelime =", *kelimePtr)
	fmt.Println("tekrar =", *tekrarPtr)
	fmt.Println("durum =", *durumPtr)
	fmt.Println("sonYazi =", sonYazi)
}

