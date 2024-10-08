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

   var t int
   fmt.Fscan(reader, &t)

   for i := 0; i < t; i++ {
      var a, b string
      fmt.Fscan(reader, &a, &b)
      count := 0
      for i := 0; i < len(a); i++ {
         if a[i] != b[i] {
            count++
         }
      }
      fmt.Fprintln(writer, fmt.Sprintf("Hamming distance is %v.", count))
   }
}
