package main

import (
   "fmt"
)

func main() {
   var t int
   fmt.Scanf("%v", &t)
   answer := ""
   for i := 0; i < t; i++ {
      var n, m int
      fmt.Scanf("%v %v", &n, &m)
      if m <4 || n <12 {
         answer += "-1\n"
         continue
      }
      sum := 11 * m + 4
      answer += fmt.Sprintf("%v\n", sum)
   }
   fmt.Println(answer)
}
