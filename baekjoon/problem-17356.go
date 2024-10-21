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

   var a, b float64
   fmt.Fscanln(reader, &a, &b)

   m := (b-a) / 400
   answer := 1/(1+math.Pow(10, m))
   
   fmt.Fprintln(writer, answer)
}
