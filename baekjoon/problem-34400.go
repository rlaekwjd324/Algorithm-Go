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
      var m int
      fmt.Fscan(reader, &m)
      
      d := m % 25
      if d < 17 {
         fmt.Fprintf(writer, "%v\n", "ONLINE")
      } else {
         fmt.Fprintf(writer, "%v\n", "OFFLINE")
      }
   }
}
