package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	sum := 0
	hasOdd := false

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		sum += x
		if x%2 == 1 {
			hasOdd = true
		}
	}

	if sum%2 == 0 {
		fmt.Println("YES")
	} else if hasOdd && n > 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
