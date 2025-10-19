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
   
   for i := 1; i <= 100; i++ {
      var d int
      fmt.Fscan(reader, &d)

      if i == 100 {
         d = d % 10
         if d == 0 {
            d = 10
         }
         fmt.Fprintf(writer, "%v\n", d)
         return
      }
   }
}
