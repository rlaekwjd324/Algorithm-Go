package main
import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Println(a*b+c*d)
}