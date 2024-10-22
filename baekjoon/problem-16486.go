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

   var d1 float64
   fmt.Fscanln(reader, &d1)
   var d2 float64
   fmt.Fscanln(reader, &d2)

   answer := d1*2.0 + d2*2.0*3.141592
   
   fmt.Fprintln(writer, answer)
}
