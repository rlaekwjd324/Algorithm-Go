package main

import (
   "fmt"
)

func main() {
   var a, b int
   fmt.Scanf("%v", &a)
   fmt.Scanf("%v", &b)
   day := a+b*7
   if day >30 {
      fmt.Println(0)
      return
   }
   fmt.Println(1)
}
