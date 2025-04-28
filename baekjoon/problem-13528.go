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
   
   var c float64
   var l int
   fmt.Fscan(reader, &c, &l)
   
   sum := 0.0
   for i := 0; i < l; i++ {
      var wi, li float64
      fmt.Fscan(reader, &wi, &li)
      
      sum += wi*li
   }
   
   fmt.Fprintf(writer, "%.7f\n", sum*c)
}
