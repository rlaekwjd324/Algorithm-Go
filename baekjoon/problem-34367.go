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

   var t string
   fmt.Fscan(reader, &t)
   
   for {
      if len(t) == 1 {
         break
      }
	   
      s := 0
      for _, v := range t {
         s += int(v-'0')
      }
	   
      t = fmt.Sprintf("%v", s)
   }
   
   fmt.Fprintf(writer, "%v\n", t)
}
