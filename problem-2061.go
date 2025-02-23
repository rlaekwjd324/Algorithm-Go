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
   
   var k, l big.Int
   fmt.Fscan(reader, &k, &l)
   
   f, i := isGoodPassword(&k, &l)
   if f {
      fmt.Fprintln(writer, "GOOD")
      return
   }
   fmt.Fprintf(writer, "BAD %v\n", i)
}

func isGoodPassword(k, l *big.Int) (bool, *big.Int) {
   for i := big.NewInt(2); i.Cmp(l) < 0; i = new(big.Int).Add(i, big.NewInt(1)) {
      m := new(big.Int).Mod(k, i)
      if m.Cmp(big.NewInt(0)) == 0 {
         return false, i
      }
   }
   return true, big.NewInt(0)
}
