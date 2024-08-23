package main

import (
   "fmt"
   "math"
)

func main() {
   var n int
   fmt.Scanln(&n)
   a := math.Sqrt(float64(n)/2.0)
   fmt.Println(a*8)
}
