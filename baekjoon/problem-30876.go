package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	var x int
	y := 1000

	for i := 0; i < n; i++ {
		pInput, _ := reader.ReadString('\n')
		pInput = strings.TrimSpace(pInput)
		nums := strings.Split(pInput, " ")

		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		if y >= b {
			x = a
			y = b
		}
	}

	fmt.Printf("%v %v\n", x, y)
}
