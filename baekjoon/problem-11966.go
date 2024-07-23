package main

import (
   "fmt"
)

func main() {
   var n float64
   fmt.Scan(&n)

   for {
      if n <=1 {
         if (n == 1) {
            fmt.Println("1")
            return
         }
         fmt.Println("0")
         return
      }
      n = n/2
   }
}
