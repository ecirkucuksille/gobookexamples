package main

import (
	"fmt"
)

func main() {
	byte_String := []byte{0x47, 0x6F, 0x6C, 0x61, 0x6E, 0x67}
	s := string(byte_String)
	fmt.Println(s)
}
