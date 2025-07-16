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

   sum := 300
   for i := 0; i < 4; i++ {
      var a int
      fmt.Fscan(reader, &a)
      
      sum += a
   }
   
   if sum <= 1800 {
      fmt.Fprintf(writer, "%v\n", "Yes")
      return
   }
   fmt.Fprintf(writer, "%v\n", "No")
}
