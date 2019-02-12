package main
import (
	"fmt"
	"net"
)
func main() {
	adres, hata := net.LookupHost("www.golang.org")
	if hata != nil {
		sonuc, ok := hata.(*net.DNSError)
		if ok {
			if sonuc.Timeout() {
				fmt.Println("timeout error")
			} else if sonuc.Temporary() {
				fmt.Println("temprorary error")
			} else {
				fmt.Println(sonuc)
			}
			return
		}
	}
	fmt.Println("adres = ", adres)
}

