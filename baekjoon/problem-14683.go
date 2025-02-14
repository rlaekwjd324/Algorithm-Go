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

   var a, b, c, d, t int
   fmt.Fscan(reader, &a, &b, &c, &d, &t)
   
   k := int(math.Abs(float64(c-a)))
   l := int(math.Abs(float64(d-b)))
   if t >= (k+l) && (t-(k+l))%2 == 0 {
      fmt.Fprint(writer, "Y")
      return
   }
   fmt.Fprint(writer, "N")
}
