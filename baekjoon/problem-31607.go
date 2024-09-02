package main

import (
   "fmt"
)

func main() {
   var a, b, c int
   fmt.Scanf("%v\n%v\n%v", &a, &b, &c)
   if a+b == c || b+c == a || c+a == b {
      fmt.Println(1)
      return
   }
   fmt.Println(0)
}
