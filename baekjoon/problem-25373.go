package main

import (
	"fmt"
	"bufio"
	"os"
   "math/big"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n big.Int
   fmt.Fscan(reader, &n)
   if n.Cmp(big.NewInt(1)) == 0 {
      fmt.Fprintf(writer, "%v\n", 1)
      return
   }
   if n.Cmp(big.NewInt(3)) <= 0 {
      fmt.Fprintf(writer, "%v\n", 2)
      return
   }
   if n.Cmp(big.NewInt(6)) <= 0 {
      fmt.Fprintf(writer, "%v\n", 3)
      return
   }
   if n.Cmp(big.NewInt(10)) <= 0 {
      fmt.Fprintf(writer, "%v\n", 4)
      return
   }
   if n.Cmp(big.NewInt(15)) <= 0 {
      fmt.Fprintf(writer, "%v\n", 5)
      return
   }
   
   num := new(big.Int).Div(&n, big.NewInt(7))
   mod := new(big.Int).Mod(&n, big.NewInt(7))
   num = new(big.Int).Add(num, big.NewInt(3))
   if mod.Cmp(big.NewInt(0)) != 0 {
      num = new(big.Int).Add(num, big.NewInt(1))
   }
   fmt.Fprintf(writer, "%v\n", num)
}
