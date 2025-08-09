package main

import (
	"fmt"
	"bufio"
	"os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var k int
   fmt.Fscan(reader, &k)
   
   for i := 0; i < k; i++ {
      var a, b int
      fmt.Fscan(reader, &a, &b)

      if a <= b {
         fmt.Fprintf(writer, "Data Set %v:\n%v\n\n", i+1, "no drought")
         continue
      }
      
      m := 0
      for a/5 > b {
         a = a/5
         m++
      }
      fmt.Fprintf(writer, "Data Set %v:\n%v%v\n\n", i+1, strings.Repeat("mega ", m), "drought")
   }
}
