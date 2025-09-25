package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n int
   var s string
   fmt.Fscan(reader, &n, &s)
   
   c := 1
   answer := ""
   for _, v := range s {
      i := int(v)+c%26
      if i > 90 {
         i -= 26
      }
      answer += string(rune(i))
      c *= 2
   }
   
   fmt.Fprintf(writer, "%v\n", answer)
}
