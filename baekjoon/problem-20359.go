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
   
   var n int
   fmt.Fscan(reader, &n)
   
   o := 0
   p := 0
   for {
      if n % 2 != 0 {
         break
      }
      n = n/2
      p++
   }
   o = n

   fmt.Fprintf(writer, "%v %v\n", o, p)
}
