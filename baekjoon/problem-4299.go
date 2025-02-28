package main

import (
   "bufio"
   "fmt"
   "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   // a+b a-b
   var sum, min int
   fmt.Fscan(reader, &sum, &min)

   if (sum+min) % 2 != 0 {
      fmt.Fprintf(writer, "%v\n", -1)
      return
   }
   a := (sum+min)/2
   b := sum - a
   if b < 0 {
      fmt.Fprintf(writer, "%v\n", -1)
      return
   }
   
   fmt.Fprintf(writer, "%v %v\n", a, b)
}
