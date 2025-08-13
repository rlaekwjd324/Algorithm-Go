package main

import (
	"fmt"
	"bufio"
	"os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

type Point struct {
   x, y int
}

func main() {
	defer writer.Flush()

   var n, m int
   fmt.Fscan(reader, &n, &m)
   
   min := math.MaxInt
   p := Point{}
   for i := 1; i <= n; i++ {
      for j := 1; j <= m; j++ {
         var x int
         fmt.Fscan(reader, &x)
         if x == 1 {
            continue
         }
         d := i + abs((m+1)/2-j)
         if min > d {
            min = d
            p = Point{i, j}
         }
      }
   }
   
   if min == math.MaxInt {
      fmt.Fprintf(writer, "%v\n", -1)
      return
   }
   fmt.Fprintf(writer, "%v %v\n", p.x, p.y)
}

func abs(num int) int {
   if num < 0 {
      return -num
   }
   return num
}
