package main

import (
   "bufio"
	"fmt"
	"os"
   "math/big"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   for i := 0; i < 3; i++ {
      var n int
      fmt.Fscan(reader, &n)
      sum := new(big.Int)
      for j := 0; j < n; j++ {
         var a big.Int
         fmt.Fscan(reader, &a)
         sum = sum.Add(sum, &a)
      }
      zero := new(big.Int)
      answer := sum.Cmp(zero)
      if answer == 1 {
         fmt.Fprintln(writer, "+")
         continue
      }
      if answer == -1 {
         fmt.Fprintln(writer, "-")
         continue
      }
      fmt.Fprintln(writer, "0")
   }

}
