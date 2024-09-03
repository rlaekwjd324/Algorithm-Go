package main

import (
   "fmt"
)

func main() {
   var a, b int
   fmt.Scanf("%v %v", &a, &b)
   count := 1
   sum := 1
   for i := 0; i < b; i++ {
      count = count + a-2
      sum += count
   }
   fmt.Println(sum)
}
