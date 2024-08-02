package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	nums := strings.Split(input, " ")
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	c, _ := strconv.Atoi(nums[2])
	if a == b && b == c {
		fmt.Println(2)
		return
	}
	if a > b {
		temp := a
		a = b
		b = temp
	}
	if b > c {
		temp := c
		c = b
		b = temp
	}
	if a*a+b*b == c*c {
		fmt.Println(1)
		return
	}
	fmt.Println(0)
}
