package main
import (
	"fmt"
	"log"
	"os"
)
var (
	fileInfo os.FileInfo
	err      error
)
func main() {
	fileInfo, err = os.Stat("/media/nazan/Data/Go.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dosya Adı:", fileInfo.Name())
	fmt.Println("Dosya Boyutu:", fileInfo.Size())
	fmt.Println("Son Değiştirme Tarihi:", fileInfo.ModTime())
}
