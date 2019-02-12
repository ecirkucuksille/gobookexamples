package main
import (
	"fmt"
	"math/rand"
	"time"
)
func indexOf(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
func main() {
	sayilar := [...]int{0, 0, 0, 0, 0, 0}

	nSource := rand.NewSource(time.Now().UnixNano())
	nRandom := rand.New(nSource)

	for i := 0; i < len(sayilar); i++ {
		for {
			sayi := nRandom.Intn(49) + 1
			if indexOf(sayilar[:], sayi) == -1 {
				sayilar[i] = sayi
				break
			}
		}
	}
	for i, v := range sayilar {
		fmt.Printf("%d. sayi = %d\n", i, v)
	}
}
