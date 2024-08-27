package main

import (
   "fmt"
)

func main() {
   answer := ""
   for {
      var t int
      fmt.Scanf("%v", &t)
      if t == -1 {
         break
      }
      sum := 0
      line := fmt.Sprintf("%v = ", t)
      for i := 1; i <= t/2; i++ {
         if t%i == 0 {
            sum += i
            line += fmt.Sprintf("%v + ", i)
         }
      }
      if sum != t {
         answer += fmt.Sprintf("%v is NOT perfect.\n", t)
      } else {
         answer += line[:len(line)-2]
         answer += "\n"
      }
   }
   fmt.Println(answer)
}
