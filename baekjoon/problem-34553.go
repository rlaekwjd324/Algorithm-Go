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

   var n string
   fmt.Fscan(reader, &n)
   
   t := int(n[0])
   s := 1
   sum := 1
   for i, v := range n {
      if i == 0 {
         continue
      }
      if int(v) > t {
         s++
      } else {
         s = 1
      }
      t = int(v)
      sum += s
   }
   
   fmt.Fprintf(writer, "%v\n", sum)
}
