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
   
   var a, b, c, t int
   fmt.Fscan(reader, &a, &b, &c, &t)
   
   sum := a
   if t > 30 {
      i := math.Ceil(float64(t-30)/float64(b))
      sum += int(i)*c
   }
   fmt.Fprintf(writer, "%v\n", sum)
}
