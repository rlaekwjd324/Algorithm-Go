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

   for {
      var t big.Int
      fmt.Fscan(reader, &t)
      zero := big.NewInt(0)
      if zero.Cmp(&t) == 0 {
         return
      }
      cal := new(big.Int)
      cal = cal.Mod(&t, big.NewInt(42))
      if zero.Cmp(cal) == 0 {
         fmt.Fprintln(writer, "PREMIADO")
      } else {
         fmt.Fprintln(writer, "TENTE NOVAMENTE")
      }
   }
}
