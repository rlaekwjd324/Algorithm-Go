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

   var a, b, c, d int
   fmt.Fscan(reader, &a, &b, &c, &d)
   
   h := a+c
   y := b+d

   if h < y {
      fmt.Fprintf(writer, "%v\n", "Hanyang Univ.")
      return
   }
   if h > y {
      fmt.Fprintf(writer, "%v\n", "Yongdap")
      return
   }
   fmt.Fprintf(writer, "%v\n", "Either")
}
