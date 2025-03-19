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
   
   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var a, b big.Int
      fmt.Fscan(reader, &a, &b)
      
      d := new(big.Int)
      m := new(big.Int)
      d = d.Div(&a, &b)
      m = m.Mod(&a, &b)
      fmt.Fprintf(writer, "%v\n%v\n\n", d, m)
   }
}
