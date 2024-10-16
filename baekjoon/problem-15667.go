package main

import (
   "bufio"
	"fmt"
	"os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n float64
   fmt.Fscanln(reader, &n)

   a := int(math.Sqrt(n))
   if a*a == int(n) {
      a--
   }
   
   fmt.Fprintln(writer, a)
}
