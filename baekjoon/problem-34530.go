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

   var d int
   fmt.Fscan(reader, &d)
   
   f := d % 360
   s := 0
   c := 0
   for {
      c++
      s += d
      if s % 360 == 0 {
         break
      }
      if s != d && s % 360 == f {
         c = -1
         break
      }
   }
   
   fmt.Fprintf(writer, "%v\n", c)
}
