package main

import (
   "fmt"
   "math"
)

func main() {
   var x, y int
   fmt.Scanf("%v", &x)
   fmt.Scanf("%v", &y)
   newX := float64(x)/2.0
   newY := float64(y)/2.0
   min := math.Min(newX, newY)
   fmt.Println(min*100)
}
