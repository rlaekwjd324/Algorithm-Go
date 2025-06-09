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

   var n, o float64
   fmt.Fscan(reader, &n, &o)
   
   x := o * n / (n-1)
   if int(o * n) % int(n-1) == 0 {
      fmt.Fprintf(writer, "%v %v\n", math.Floor(x)-1, math.Ceil(x))
      return
   }
   fmt.Fprintf(writer, "%v %v\n", math.Floor(x), math.Floor(x))
}
