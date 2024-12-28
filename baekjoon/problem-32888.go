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

   var a, b float64
   fmt.Fscan(reader, &a, &b)

    var c float64
   c = a*a + b*b
  
   fmt.Fprintf(writer, "%v\n", math.Sqrt(c))
}
