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

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var lt, wt, le, we int
      fmt.Fscan(reader, &lt, &wt, &le, &we)
      
      if lt*wt < le*we {
         fmt.Fprintf(writer, "%v\n", "Eurecom")
         continue
      }
      if lt*wt > le*we {
         fmt.Fprintf(writer, "%v\n", "TelecomParisTech")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "Tie")
   }
}
