package main

import (
   "fmt"
)

func main() {
   var n int
   fmt.Scanln(&n)
   count := 0
   for {
      if n == 1 {
         break
      }
      if n%2 == 1 {
         n = n+1
         count++
      }
      n = n/2
      count++
   }
   fmt.Println(count)
}
