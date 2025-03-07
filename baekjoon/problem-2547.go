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
      var n int
      fmt.Fscan(reader, &n)
      sum := big.NewInt(0)
      for j := 0; j < n; j++ {
         var a big.Int
         fmt.Fscan(reader, &a)
         sum.Add(sum, &a)
      }
      
      result := new(big.Int)
      result.Mod(sum, big.NewInt(int64(n)))
      if result.Cmp(big.NewInt(0)) == 0 {
         fmt.Fprintln(writer, "YES")
      }else {
         fmt.Fprintln(writer, "NO")
      }
   }
}
