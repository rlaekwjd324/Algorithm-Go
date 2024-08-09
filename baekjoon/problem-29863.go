package main

import (
   "fmt"
)

func main() {
   var x, y int
   fmt.Scanf("%v", &x)
   fmt.Scanf("%v", &y)
   
   sum := 0
   if x >= 20 {
      sum = (24-x)
      sum += y
   }else {
      sum = y-x
   }
   fmt.Println(sum)
}
