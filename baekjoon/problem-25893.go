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
   fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
   answer := ""
   for i := 0; i < n; i++ {
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      nums := strings.Split(input, " ")
      a, _ := strconv.Atoi(nums[0])
      b, _ := strconv.Atoi(nums[1])
      c, _ := strconv.Atoi(nums[2])
      count := 0
      if a >= 10 {
         count++
      }
      if b >= 10 {
         count++
      }
      if c >= 10 {
         count++
      }
      print := ""
      if count == 3 {
         print = "triple-double"
      } else if count == 2 {
         print = "double-double"
      }else if count == 1 {
         print = "double"
      }else if count == 0 {
         print = "zilch"
      }
      answer += fmt.Sprintf("%v\n%v\n\n", input, print)
   } 
   fmt.Println(answer)
}
