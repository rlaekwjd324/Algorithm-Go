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
   
   for {
      var a, b int
      fmt.Fscan(reader, &a, &b)

      if a == 0 && b == 0 {
         break
      }

      c := a / b
      d := a % b
      
      fmt.Fprintf(writer, "%v %v / %v\n", c, d, b)
   }
}
