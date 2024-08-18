package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	sum := n + m
	str := strconv.Itoa(sum)
	fmt.Println(len(str))
}
