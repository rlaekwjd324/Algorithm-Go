package main

import (
   "fmt"
)

func main() {
   var s, a, b int
   fmt.Scan(&s)
   fmt.Scan(&a)
   fmt.Scan(&b)

   price := 250
   height := a

   for {
      if height >= s {
         fmt.Println(price)
         return
      }

      height += b
      price += 100
   }

 }
