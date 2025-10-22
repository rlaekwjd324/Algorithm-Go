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
   
   var a, b, n, k int
   fmt.Fscan(reader, &a, &b, &n, &k)
   
   c := 0
   ta := 1
   tb := 1
   tn := 0
   for {
      if c == k {
         break
      }
      c++
      tn++
      if tn >= n {
         tb++
         tn = 0
         if tb > b {
            ta++
            tb = 1
            if ta > a {
               ta = 1
            }
         }
      }
   }

   fmt.Fprintf(writer, "%v %v\n", ta, tb)
}
