package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
)

func main() {
   var n int
   fmt.Scanf("%v", &n)
   reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   var max = 0
   for i := 0; i < n; i++ {
      num, _ := strconv.Atoi(nums[i])
      if max < num-n+i {
         max = num-n+i
      }
   }
   fmt.Println(max)
 }
