package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var answer int = 1
	for i := 1; i <= n; i++ {
		answer *= i
	}
	fmt.Println(answer)
}