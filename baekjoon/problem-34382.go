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

   var t int
   fmt.Fscan(reader, &t)
   
   arr := make([]float64, t)
   s := 0.0
   for i := 0; i < t; i++ {
      fmt.Fscan(reader, &arr[i])
      s += arr[i]
   }
   c := 0
   a := s/float64(t)
   for i := 0; i < t; i++ {
      if math.Abs(arr[i]-a) > 10 {
         c++
      }
   }
   fmt.Fprintf(writer, "%v\n", c)
}
