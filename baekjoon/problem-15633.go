package main

import (
   "fmt"
   "math"
)

func main() {
   var n int
   fmt.Scanf("%v", &n)
   sum := 0
   for i := 1; float64(i) <= math.Sqrt(float64(n)); i++ {
      if n % i == 0 {
         sum += i
         sum += (n/i)
      }
      if math.Sqrt(float64(n)) == float64(i) {
         sum -= i
      }
   }
   fmt.Println(sum*5-24)
}
