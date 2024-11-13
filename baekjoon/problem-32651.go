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

   var n int
   fmt.Fscanln(reader, &n)
   
   if n >= 100000 {
      fmt.Fprintf(writer, "No")
      return
   }
   if n % 2024 != 0 {
      fmt.Fprintf(writer, "No")
      return
   }

   fmt.Fprintf(writer, "Yes")
}
