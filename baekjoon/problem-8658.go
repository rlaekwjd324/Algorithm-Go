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
   
   min := 1000000000
   for i := 2; i < n; i++ {
      if n % i != 0 {
         min = i
         break
      }
   }
   max := 0
   for i := n-1; i >= 2; i-- {
      if n % i != 0 {
         max = i
         break
      }
   }

   fmt.Fprintf(writer, "%v %v\n", min, max)
}
