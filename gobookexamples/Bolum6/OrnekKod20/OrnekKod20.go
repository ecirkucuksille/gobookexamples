package main
import (
	"fmt"
	"strconv"
)   
func main() {
	b, _ := strconv.ParseBool("true") // string - bool dönüşümü 
	f, _ := strconv.ParseFloat("21.54", 64) // string - float dönüşümü
	i, _ := strconv.ParseInt("16", 10, 64) // string - int dönüşümü
	u, _ := strconv.ParseUint("16", 16,64) // string - uint dönüşümü
	fmt.Printf("%T, %v \n",b,b)
	fmt.Printf("%T, %v \n",f,f)
	fmt.Printf("%T, %v \n",i,i)
	fmt.Printf("%T, %v \n",u,u)
	}