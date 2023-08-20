package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		var num string
		fmt.Scanln(&num)
		var s = strings.Split(num, "")
		b, _ := strconv.Atoi(s[len(num)-1])
		if(b%2 == 0) {
			fmt.Println("even")
		}else {
			fmt.Println("odd")
		}
	}
}