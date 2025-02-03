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

   var d, h float64 
   fmt.Fscan(reader, &d, &h)

   r := d/2 + 5
   
   side := math.Pi * r * 2 * h
   top := math.Pi * r * r
   
   fmt.Fprintln(writer, side+top)
}
