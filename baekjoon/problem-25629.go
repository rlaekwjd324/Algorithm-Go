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
   fmt.Scan(&n)

   reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")

   evenCount := 0
   oddCount := 0

   for i := 0; i < n; i++ {
      num, _ := strconv.Atoi(nums[i])
      if num % 2 == 0 {
         evenCount++
         continue
      }
      oddCount++
   }

   if evenCount > oddCount || evenCount + 1 < oddCount {
      fmt.Println(0)
      return
   }
   fmt.Println(1)
 }
