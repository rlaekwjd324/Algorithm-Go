package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, n int
	fmt.Scanln(&x)
	fmt.Scanln(&n)
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	nums := strings.Split(input, " ")
	sum := float64(0)
	if n%2 == 0 {
		for i := 0; i < x; i++ {
			num, _ := strconv.Atoi(nums[i])
			sum += math.Pow(float64(num), float64(n))
		}
	} else {
		for i := 0; i < x; i++ {
			num, _ := strconv.Atoi(nums[i])
			if num < 0 {
				continue
			}
			sum += math.Pow(float64(num), float64(n))
		}
	}
	fmt.Println(sum)
}
