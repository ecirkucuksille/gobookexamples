package main

import (
	"fmt"
)

func main() {
	fmt.Print(" 1 ")
	goto End
	fmt.Print(" 2 ")
End:
	fmt.Print(" 3 ")
}
