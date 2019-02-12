package main
import (  
    "fmt"
    "os"
)
func main() {  
    f, err := os.Create("/media/nazan/Data/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(" Golang ")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "byte dosyaya yazdırıldı")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}
