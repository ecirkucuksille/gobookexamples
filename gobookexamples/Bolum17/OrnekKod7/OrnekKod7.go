package main
import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var jsonVeri map[string]interface{}
	dosya, _ := os.Open("ornek.json")
	json.NewDecoder(dosya).Decode(&jsonVeri)
	fmt.Println(jsonVeri)

}
