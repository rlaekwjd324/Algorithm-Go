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
   
   var z int
   fmt.Fscan(reader, &z)
   
   s := "ABCBCDCDADAB"
   for i := 0; i < z; i++ {
      var n int
      fmt.Fscan(reader, &n)
      
      d := (n-1) % 12

      fmt.Fprintf(writer, "%v\n", string(s[d]))
   }
}
