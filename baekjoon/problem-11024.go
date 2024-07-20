package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := ""
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	x, _ := strconv.Atoi(input)
	for i := 1; i <= x; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		nums := strings.Split(input, " ")
		sum := 0
		for j := 0; j < len(nums); j++ {
			num, _ := strconv.Atoi(nums[j])
			sum += num
		}
		answer += fmt.Sprintf("%v\n", sum)
	}
	fmt.Println(answer)
}
