package main

import (
	"fmt"
	"bufio"
	"os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n int
   var x, y float64
   fmt.Fscan(reader, &n, &x, &y)
   
   for i := 0; i < n-1; i++ {
      var dir string
      var d float64
      fmt.Fscan(reader, &dir, &d)
      
      dd := math.Sqrt(d*d/2)
      switch dir {
      case "N":
         y += d
      case "S":
         y -= d
      case "E":
         x += d
      case "W":
         x -= d
      case "NE":
         x += dd
         y += dd
      case "SE":
         x += dd
         y -= dd
      case "NW":
         x -= dd
         y += dd
      case "SW":
         x -= dd
         y -= dd
      }
   }
   fmt.Fprintf(writer, "%.8f %.8f\n", x, y)
}
