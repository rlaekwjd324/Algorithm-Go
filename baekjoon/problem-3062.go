package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)

	for i := 0; i < n; i++ {
		var num string
		fmt.Scanf("%v", &num)
		reverseNum := Reverse(num)
		numInt, _ := strconv.Atoi(num)
		reverseNumInt, _ := strconv.Atoi(reverseNum)
		sum := numInt + reverseNumInt
		if isSameLR(strconv.Itoa(sum)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isSameLR(sum string) bool {
	temp := strings.Split(sum, "")
	for i := 0; i < len(sum)/2; i++ {
		if temp[i] != temp[len(sum)-i-1] {
			return false
		}
	}
	return true
}
