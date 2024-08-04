package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)
	if n%4 == 0 {
		fmt.Println("Even")
	} else if n%2 == 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Either")
	}
}
