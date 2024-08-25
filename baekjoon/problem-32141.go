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
   h, _ := strconv.Atoi(nums[1])
   input, _ = reader.ReadString('\n')
   input = strings.TrimSpace(input)
   nums = strings.Split(input, " ")
   sum := 0
   count := 0
   for i := 0; i < n; i++ {
      num, _ := strconv.Atoi(nums[i])
      sum += num
      count += 1
      if sum >= h {
         break
      }
   }
   if sum < h {
      fmt.Println(-1)
      return
   }
   fmt.Println(count)
}
