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

   for i := 101; i < 1000; i++ {
      for j := 101; j < 1000; j++ {
         if i == j {
            continue
         }
         if i%10 != j/100 {
            continue
         }
         if i*(j%100) == j*(i/10) {
            fmt.Fprintf(writer, "%v / %v = %v / %v\n", i, j, (i/10), (j%100))
         }
      }
   }
}
