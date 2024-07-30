package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
)

func main() {
   reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   n, _ := strconv.Atoi(nums[0])
   m, _ := strconv.Atoi(nums[1])
   min := n
   if (n > m) {
      min = m
   }
   answer := min/2
   fmt.Println(answer)
 }
