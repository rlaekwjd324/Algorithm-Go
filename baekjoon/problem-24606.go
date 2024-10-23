package main

import (
   "bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var p1 string
   fmt.Fscanln(reader, &p1)
   var p2 string
   fmt.Fscanln(reader, &p2)

   count := 1
   for i := 0; i < len(p1); i++ {
      if p1[i] != p2[i] {
         count = count * 2
      }
   }
   
   fmt.Fprintln(writer, count)
}
