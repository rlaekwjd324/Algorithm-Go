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

   var n float64
   var m int
   fmt.Fscan(reader, &n, &m)
   
   for i := 0; i < m; i++ {
      var a float64
      fmt.Fscan(reader, &a)
      
      if math.Abs(a-n) >= math.Abs(a-1) {
         fmt.Fprintf(writer, "%v ", 1)
      } else {     
         fmt.Fprintf(writer, "%v ", n)
      }
   }
}
