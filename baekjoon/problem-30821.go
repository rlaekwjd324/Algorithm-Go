package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	answer := 1
	for i := 1; i <= 5; i++ {
		answer = answer * (n - i + 1)
		answer = answer / i
	}
	fmt.Println(answer)
}
