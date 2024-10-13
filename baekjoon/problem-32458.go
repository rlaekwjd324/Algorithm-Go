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
   fmt.Fscan(reader, &n)
  
   m := math.Floor(n)
   fmt.Fprintln(writer, int(m))
}
