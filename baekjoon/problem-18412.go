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

   var s, a, b int
   var w string
   fmt.Fscan(reader, &s, &a, &b, &w)
   
   n := w[:a-1]
   r := w[a-1:b]
   n += Reverse(r)
   n += w[b:s]

   fmt.Fprintf(writer, "%s\n", n)

}

func Reverse(s string) (result string) {
   for _,v := range s {
     result = string(v) + result
   }
   return 
 }
