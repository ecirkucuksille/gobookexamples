package main
import (
	"fmt"
	"strconv"
)   
func main() {
	v := "10"
	fmt.Printf("%T , %v \n", v,v)
	sayi , _ := strconv.Atoi(v)
	fmt.Printf("%T, %v \n", sayi, sayi)
	}