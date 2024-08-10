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
   k, _ := strconv.Atoi(nums[2])
   min := n-m*k
   max := n-m*(k-1)-1
   if min < 0 {
      min = 0
   }
   if max < 0 {
      max = 0
   }
   fmt.Printf("%v %v\n", min, max)
 }
