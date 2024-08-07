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
   a, _ := strconv.Atoi(nums[0])
   b, _ := strconv.Atoi(nums[1])
   num := 1
   for i := a; i <= b; i++ {
      sum := 0
      for j := 1; j <= i; j++ {
         sum += j
      }
      num *= sum
      num = num%14579
   }
   fmt.Println(num)
 }
