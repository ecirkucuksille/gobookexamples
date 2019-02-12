package main
import (  
    "fmt"
    "io/ioutil"
)
func main() {  
    data, err := ioutil.ReadFile("/media/nazan/Data/deneme.txt")
    if err != nil {
        fmt.Println("Dosya okuma hatası", err)
        return
    }
	fmt.Println(string(data))
}
