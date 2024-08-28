package main

import (
   "fmt"
   "math"
)

func main() {
   var a int
   fmt.Scanf("%v", &a)
   temp := (math.Sqrt(float64(a))*4)
   if int(temp*100000000)%100000000 == 0 {
      fmt.Printf("%.1f\n", temp)
      return
   }
   fmt.Printf("%.8f\n", temp)
}
