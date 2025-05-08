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
   
   var n float64
   fmt.Fscan(reader, &n)
   
   s := 0.0
   for i := 0; i < int(n); i++ {
      var a float64
      fmt.Fscan(reader, &a)
      
      s += a * a * a
   }

   b := math.Pow(s, 1.0/3.0)
   
   fmt.Fprintf(writer, "%v\n", b)
}
