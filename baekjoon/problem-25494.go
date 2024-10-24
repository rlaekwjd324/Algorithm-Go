package main

import (
   "bufio"
	"fmt"
	"os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var t int
   fmt.Fscanln(reader, &t)
   
   answer := ""

   for i := 0; i < t; i++ {
      var a, b, c float64
      fmt.Fscanln(reader, &a, &b, &c)
      min := math.Min(a, b)
      min = math.Min(min, c)
      answer += fmt.Sprintf("%v\n", int(min))
   }
   
   fmt.Fprintln(writer, answer)
}
