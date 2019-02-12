package main
import (
	"fmt"
	"strings"
)
func main() {
	s:= "Golang,Programlama"
	res := strings.Split(s,",")
              for i := range res {
                      fmt.Println(res[i])
             }
             fmt.Println(len(res))
}